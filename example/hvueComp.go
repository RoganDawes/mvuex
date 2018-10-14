// +build js

package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/HuckRidgeSW/hvue"
)

func InitHvueComp() {
	hvue.NewComponent(
		"hvuecomp",
		hvue.Template(compTemplate),

		// map Vuex store.state.count to computed property
		hvue.Computed("count", func(vm *hvue.VM) interface{} {
			return vm.Get("$store").Get("state").Get("count")
		}),

		// map Vuex store.state.text to computed property "text", use mutation "setText" as setter
		hvue.ComputedWithGetSet("text",
			func(vm *hvue.VM) interface{} {
				return vm.Get("$store").Get("state").Get("text")
			},
			func(vm *hvue.VM, newValue *js.Object) {
				vm.Get("$store").Call("commit", "setText", newValue) //Send back changes to Vuex state via committed mutation
			}),
		// map Vuex Action dispatch of "actiontest" this.$store.dispatch("actiontest") to Vue method "actionincrement"
		hvue.Method("actionincrement", func(vm *hvue.VM, count *js.Object) {
			vm.Get("$store").Call("dispatch", "actiontest")
		}),
		// map Vuex Mutation commit of "increment" with argument this.$store.commit("increment",count) to Vue method "increment(count)"
		hvue.Method("increment", func(vm *hvue.VM, count *js.Object) {
			// normal way to access the store.commit() function
			vm.Get("$store").Call("commit", "increment", count)
		}),
		hvue.Method("decrement", func(vm *hvue.VM) {
			// normal way to access the store.commit() function
			vm.Get("$store").Call("commit", "decrement")
		}),
	)
}

const compTemplate = `
<div>
  <input v-model="text"></input>
  <p>{{ count }}</p>
  <p>
    <button @click="increment(1,2,3)">+</button>
	<button @click="increment(2)">+2</button>
	<button @click="actionincrement">Action</button>
    <button @click="decrement">-</button>
  </p>
</div>
`

