package main

import (
	"github.com/HuckRidgeSW/hvue"
	"github.com/mame82/mvuex"
	"github.com/gopherjs/gopherjs/js"
)

func Store(store *mvuex.Store) hvue.ComponentOption {
	return func(config *hvue.Config) {
		config.Set("store", store)
	}
}

func main() {
	store := NewMVuexStore()
	InitHvueComp() //Register a Vue.js component using the Vux store

	vm := hvue.NewVM(
		hvue.El("#app"),
		hvue.Template(appTemplate),
		Store(store), //Add store to VM
	)

	js.Global.Set("vm", vm) //expose vm for debugging
}

const appTemplate = `
<div>
<h1>Vuex bindings in combined with Vue bindings (hvue)</h1>
<p>Two instance of the same Vue component share the same Vuex store</p>
<table>
  <tr>
    <th>Instance 1</th>
    <th>Instance 2</th>
  </tr>
  <tr>
    <td style='background: lightcoral'><hvuecomp></hvuecomp></td>
    <td style='background: lightblue'><hvuecomp></hvuecomp></td>
  </tr>
</table> 
</div>
`
