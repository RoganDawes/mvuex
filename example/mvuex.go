// +build js

package main

import (
	"github.com/gopherjs/gopherjs/js"
	"time"
	"github.com/mame82/mvuex"
)


func O() *js.Object { return js.Global.Get("Object").New() }

type StoreState struct {
	*js.Object
	Counter int `js:"count"`
	Text string `js:"text"`
}

func createStoreStateStruct() StoreState {
	state := StoreState{Object:O()}
	state.Counter = 1337
	state.Text = "Hi there says MaMe82"
	return state
}



func NewMVuexStore() *mvuex.Store {
	state := createStoreStateStruct()
	store := mvuex.NewStore(
		mvuex.State(state), // add state to Vuex store
		mvuex.Action("actiontest", func(store *mvuex.Store, context *mvuex.ActionContext, state *StoreState) {
			// Intended to run asynchronous, but has to be wrapped in go routine anyways
			go func() {
				for i:=0; i<10; i++ {
					println(state.Counter)
					time.Sleep(1*time.Second)
					context.Commit("increment",5) //commit the mutation "increment" with argument 5 to the action context
				}

			}()

		}),
		mvuex.Mutation("increment", func (store *mvuex.Store, state *StoreState, add int) {
			// Mutation accepting an int argument, increments the counter
			state.Counter += add
			return
		}),
		mvuex.Mutation("decrement", func (store *mvuex.Store, state *StoreState) {
			// Mutation without argument, decrements the counter
			state.Counter--
			return
		}),
		mvuex.Mutation("setText", func (store *mvuex.Store, state *StoreState, newText string) {
			state.Text = newText
			return
		}),
	)

	return store
	// propagate Vuex store to global scope to allow injecting it to Vue by setting the "store" option
	//js.Global.Set("store", store)
}
