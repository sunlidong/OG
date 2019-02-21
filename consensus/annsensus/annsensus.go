package annsensus

import (
	"fmt"
	"sync"
	"time"

	"github.com/annchain/OG/common/crypto"
	"github.com/annchain/OG/og"
	"github.com/annchain/OG/types"
	log "github.com/sirupsen/logrus"
)

type AnnSensus struct {
	doCamp bool // the switch of whether annsensus should produce campaign.

	newTxHandlers []chan types.Txi

	campaignFlag bool
	maxCamps     int
	candidates   map[types.Address]*types.Campaign
	alsorans     map[types.Address]*types.Campaign
	campaigns    map[types.Address]*types.Campaign // replaced by candidates

	termchgFlag bool

	campsCh   chan []*types.Campaign
	termchgCh chan *types.TermChange

	Hub            *og.Hub //todo use interface later
	Txpool         og.ITxPool
	Idag           og.IDag
	partner        *Partner
	MyPrivKey      *crypto.PrivateKey
	Threshold      int
	NbParticipants int
	startGossip    chan bool

	mu          sync.RWMutex
	termchgLock sync.RWMutex
	close       chan struct{}
}

func NewAnnSensus(campaign bool) *AnnSensus {
	return &AnnSensus{
		close:         make(chan struct{}),
		newTxHandlers: []chan types.Txi{},
		campaignFlag:  campaign,
		campaigns:     make(map[types.Address]*types.Campaign),
		startGossip:   make(chan bool),
	}
}

func (as *AnnSensus) Start() {
	log.Info("AnnSensus Start")
	if as.campaignFlag {
		as.ProdCampaignOn()
		// TODO campaign gossip starts here?
		go as.gossipLoop()
	}
	// TODO
}

func (as *AnnSensus) Stop() {
	log.Info("AnnSensus Stop")
	as.ProdCampaignOff()
	close(as.close)
}

func (as *AnnSensus) Name() string {
	return "AnnSensus"
}

func (as *AnnSensus) GetBenchmarks() map[string]interface{} {
	// TODO
	return nil
}

// RegisterReceiver add a channel into AnnSensus.newTxHandlers, newTxHandlers
// are the hooks that handle new txs.
func (as *AnnSensus) RegisterReceiver(c chan types.Txi) {
	as.newTxHandlers = append(as.newTxHandlers, c)
}

// ProdCampaignOn let annsensus start producing campaign.
func (as *AnnSensus) ProdCampaignOn() {
	as.doCamp = true
	go as.prodcampaign()
}

// ProdCampaignOff let annsensus stop producing campaign.
func (as *AnnSensus) ProdCampaignOff() {
	as.doCamp = false
}

// campaign continuously generate camp tx until AnnSensus.CampaingnOff is called.
func (as *AnnSensus) prodcampaign() {
	// TODO
	for {
		select {
		case <-as.close:
			log.Info("campaign stopped")
			return
		case <-time.After(time.Second * 10):
			if !as.doCamp {
				log.Info("campaign stopped")
				return
			}
			// generate campaign.
			camp := as.genCamp()
			// send camp
			if camp != nil {
				for _, c := range as.newTxHandlers {
					c <- camp
				}
			}
		}
	}
}

// commit takes a list of campaigns as input and record these
// camps' information It checks if the number of camps reaches
// the threshold. If so, start term changing flow.
func (as *AnnSensus) commit(camps []*types.Campaign) {
	// TODO

	for i, c := range camps {
		if as.isTermChanging() {
			// add those unsuccessful camps into alsoran list.
			as.addAlsorans(camps[i:])
			return
		}
		// TODO
		// handle campaigns should not only add it into candidate list.
		as.candidates[c.Issuer] = c
		if as.canChangeTerm() {
			as.termchgLock.Lock()
			as.termchgFlag = true
			as.termchgLock.Unlock()
			go as.changeTerm()
		}

	}

}

// canChangeTerm returns true if the campaigns cached reaches the
// term change requirments.
func (as *AnnSensus) canChangeTerm() bool {
	// TODO
	return true
}

func (as *AnnSensus) isTermChanging() bool {
	as.termchgLock.RLock()
	changing := as.termchgFlag
	as.termchgLock.RUnlock()

	return changing
}

func (as *AnnSensus) changeTerm() {
	// TODO
}

// addAlsorans add a list of campaigns into alsoran list.
func (as *AnnSensus) addAlsorans(camps []*types.Campaign) {
	// TODO
}

func (as *AnnSensus) loop() {

	for {
		select {
		case <-as.close:
			return

		case camps := <-as.campsCh:
			fmt.Println(camps)
			// TODO
			// case commit

		case termchg := <-as.termchgCh:
			fmt.Println(termchg)
			// TODO
			// case start term change gossip
		}
	}

}
