package main

import (
	"fmt"
	"math/rand"
	"time"
)

type CardSuit string

const (
	CLUBS    CardSuit = "clubs"
	DIAMONDS CardSuit = "diamonds"
	HEARTS   CardSuit = "hearts"
	SPADES   CardSuit = "spades"
)

type Card struct {
	suit  CardSuit
	value int
}

func NewCard(suit CardSuit, value int) *Card {
	return &Card{
		suit:  suit,
		value: value,
	}
}

func (c *Card) GetSuit() CardSuit {
	return c.suit
}

func (c *Card) GetValue() int {
	return c.value
}

func (c *Card) Print() {
	fmt.Println(c.GetSuit(), c.GetValue())
}

type Deck struct {
	cards []*Card
}

func NewDeck() *Deck {
	deck := &Deck{
		cards: []*Card{},
	}
	for _, suit := range []CardSuit{CLUBS, DIAMONDS, HEARTS, SPADES} {
		for value := 1; value <= 13; value++ {
			deck.cards = append(deck.cards, NewCard(suit, min(value, 10)))
		}
	}
	return deck
}

func (d *Deck) Print() {
	for _, card := range d.cards {
		card.Print()
	}
}

func (d *Deck) Draw() *Card {
	if len(d.cards) == 0 {
		return nil
	}
	card := d.cards[len(d.cards)-1]
	d.cards = d.cards[:len(d.cards)-1]
	return card
}

func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	for i := range d.cards {
		j := rand.Intn(len(d.cards))
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	}
}

type Hand struct {
	cards []*Card
	score int
}

func NewHand() *Hand {
	return &Hand{
		cards: []*Card{},
		score: 0,
	}
}

func (h *Hand) AddCard(card *Card) {
	h.cards = append(h.cards, card)
	if card.GetValue() == 1 {
		h.score += 11
		if h.score+11 > 21 {
			h.score++
		}
	} else {
		h.score += card.GetValue()
	}
}

func (h *Hand) GetScore() int {
	return h.score
}

func (h *Hand) GetCards() []*Card {
	return h.cards
}

func (h *Hand) Print() {
	for _, card := range h.cards {
		fmt.Println(card.GetSuit(), card.GetValue())
	}
}

type PlayerN interface {
	GetHand() *Hand
	ClearHand()
	AddCard(card *Card)
	MakeMove() bool
}

type UserPlayer struct {
	balance int
	hand    *Hand
}

func NewUserPlayer(balance int, hand *Hand) *UserPlayer {
	return &UserPlayer{
		balance: balance,
		hand:    hand,
	}
}

func (p *UserPlayer) GetHand() *Hand {
	return p.hand
}

func (p *UserPlayer) ClearHand() {
	p.hand = NewHand()
}

func (p *UserPlayer) AddCard(card *Card) {
	p.hand.AddCard(card)
}

func (p *UserPlayer) PlaceBet(amount int) int {
	if amount > p.balance {
		panic("Insufficient funds")
	}
	p.balance -= amount
	return amount
}

func (p *UserPlayer) ReceiveWinnings(amount int) {
	p.balance += amount
}

func (p *UserPlayer) GetBalance() int {
	return p.balance
}

func (p *UserPlayer) MakeMove() bool {
	if p.hand.GetScore() > 21 {
		return false
	}
	var move string
	fmt.Print("Draw card? [y/n] ")
	fmt.Scanln(&move)
	return move == "y"
}

type Dealer struct {
	hand        *Hand
	targetScore int
}

func NewDealer(hand *Hand) *Dealer {
	return &Dealer{
		hand:        hand,
		targetScore: 17,
	}
}

func (d *Dealer) GetHand() *Hand {
	return d.hand
}

func (d *Dealer) ClearHand() {
	d.hand = NewHand()
}

func (d *Dealer) AddCard(card *Card) {
	d.hand.AddCard(card)
}

func (d *Dealer) MakeMove() bool {
	return d.hand.GetScore() < d.targetScore
}

func (d *Dealer) UpdateTargetScore(score int) {
	d.targetScore = score
}

type GameRound struct {
	player *UserPlayer
	dealer *Dealer
	deck   *Deck
}

func NewGameRound(player *UserPlayer, dealer *Dealer, deck *Deck) *GameRound {
	return &GameRound{
		player: player,
		dealer: dealer,
		deck:   deck,
	}
}

func (g *GameRound) getBetUser() int {
	var amount int
	fmt.Print("Enter a bet amount: ")
	fmt.Scanln(&amount)
	return amount
}

func (g *GameRound) dealInitialCards() {
	for i := 0; i < 2; i++ {
		g.player.AddCard(g.deck.Draw())
		g.dealer.AddCard(g.deck.Draw())
	}
	fmt.Println("Player hand:")
	g.player.GetHand().Print()
	dealerCard := g.dealer.GetHand().GetCards()[0]
	fmt.Println("Dealer's first card:")
	dealerCard.Print()
}

func (g *GameRound) cleanupRound() {
	g.player.ClearHand()
	g.dealer.ClearHand()
	fmt.Println("Player balance:", g.player.GetBalance())
}

func (g *GameRound) Play() {
	g.deck.Shuffle()

	if g.player.GetBalance() <= 0 {
		fmt.Println("Player has no more money =)")
		return
	}
	userBet := g.getBetUser()
	g.player.PlaceBet(userBet)

	g.dealInitialCards()

	for g.player.MakeMove() {
		drawnCard := g.deck.Draw()
		fmt.Println("Player draws", drawnCard.GetSuit(), drawnCard.GetValue())
		g.player.AddCard(drawnCard)
		fmt.Println("Player score", g.player.GetHand().GetScore())
	}
	if g.player.GetHand().GetScore() > 21 {
		fmt.Println("Player loses")
		g.cleanupRound()
		return
	}

	playerScore := g.player.GetHand().GetScore()
	g.dealer.UpdateTargetScore(playerScore)
	for g.dealer.MakeMove() {
		g.dealer.AddCard(g.deck.Draw())
	}

	dealerScore := g.dealer.GetHand().GetScore()
	if dealerScore > 21 || playerScore > dealerScore {
		fmt.Println("Player wins")
		g.player.ReceiveWinnings(userBet * 2)
	} else if dealerScore > playerScore {
		fmt.Println("Player loses")
	} else {
		fmt.Println("Game ends in a draw")
		g.player.ReceiveWinnings(userBet)
	}
	g.cleanupRound()
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func mainB() {
	player := NewUserPlayer(1000, NewHand())
	dealer := NewDealer(NewHand())

	for player.GetBalance() > 0 {
		deck := NewDeck()
		round := NewGameRound(player, dealer, deck)
		round.Play()
	}
}
