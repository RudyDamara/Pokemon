import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'
import VueAxios from 'vue-axios'


Vue.use(Vuex)
Vue.use(VueAxios, axios)

export default new Vuex.Store({
  state: {
    pokemonlists: null,
    loading: true,
    catchPokemon:null,
    myPokemon:null,
    release:null,
    update:null,
 
  },
  mutations: {
    LOAD_SUCCESS(state, posts) {
      state.pokemonlists = posts
    },
    CATCH_SUCCESS(state,posts){
      state.catchPokemon = posts
    },
    MY_POKEMON(state, posts){
      state.myPokemon =  posts
    },
    RELEASE_SUCCESS(state, posts){
      state.release =  posts
    },

    UPDATE_SUCCESS(state, posts){
      state.update =  posts
    },
  },
  getters: {
      items: state => {
          return state.pokemonlists;
      },
      responseCatchPokemon:state=>{
        return state.catchPokemon
      },
      responseMyPokemon:state=>{
        return state.myPokemon
      },
      release:state=>{
        return state.release
      },
      update:state=>{
        return state.update
      }
  },
  actions: {
    loadData ({ commit }) {
      axios
        .get('https://pokeapi.co/api/v2/pokemon?limit=10&offset=0')
        .then(response =>  {
           var data = []
            for (const iterator of response.data.results) {
              const element = {}
              element.name = iterator.name
              element.url = iterator.url
              axios
              .get(iterator.url)
              .then(resp =>  {
                  element.detail = resp.data
                  element.gambar1 = resp.data.sprites.front_shiny
                  element.gambar = "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/"+resp.data.id+".png"
              })
              data.push(element)
            }
          commit('LOAD_SUCCESS', data)
        })
    },
    loadCatchData ({ commit }) {
      axios
        .get('http://localhost:1323/pokemon')
        .then(response =>  {
          console.log(response)
          var data = []
            for (const iterator of response.data.data) {
              const element = {}
              element.name = iterator.name
              element.name_pokemon = iterator.name_pokemon
              element.url = iterator.url_pokemon 
              element.id = iterator.id
              axios
              .get(iterator.url_pokemon)
              .then(resp =>  {
                  element.detail = resp.data
                  element.gambar1 = resp.data.sprites.front_shiny
                  element.gambar = "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/"+resp.data.id+".png"
              })

              data.push(element)
            }
            commit('MY_POKEMON', data)
        })
    },
    saveData ({ commit }, payload) {
      axios
        .post('http://localhost:1323/pokemon/catch', payload)
        .then(response =>  {
          if(response.data.code == 200){
            commit('CATCH_SUCCESS', true)
          }else{
            commit('CATCH_SUCCESS', false)
          }
        })
    },
    realase({ commit }, payload) {
      axios
        .delete('http://localhost:1323/pokemon/'+payload.id)
        .then(response =>  {
          if(response.data.code == 200){
            commit('RELEASE_SUCCESS', true)
          }else{
            commit('RELEASE_SUCCESS', false)
          }
        })
    },


    update({ commit }, payload) {
      var data = {
        name: payload.name,
      }
      axios
        .put('http://localhost:1323/pokemon/'+payload.id, data)
        .then(response =>  {
          if(response.data.code == 200){
            commit('UPDATE_SUCCESS', true)
          }else{
            commit('UPDATE_SUCCESS', false)
          }
        })
    },

  },
  modules: {
  }
})
