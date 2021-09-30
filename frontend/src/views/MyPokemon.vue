<template>
  <div>
    <MyPokemonC :item=doneData />
    </div>
</template>

<script>

  import MyPokemonC from '../components/MyPokemon'
  export default {
    name: 'MyPokemon',

    components: {
      MyPokemonC,
    },
    computed:{
      doneData () {
        return this.$store.getters.responseMyPokemon
      }
    },
    methods: {
      // clickDetail(param) {
      //   this.$router.push({
      //     name: "RenamePokemon",
      //     params: {
      //       data: param,
      //     },
      //   });
      // },
    },
    watch:{
      doneData(){
        this.itemPokemon= this.doneData
      }
    },  
    data() {
      return {
        data:null,
        }
    },
    mounted() {
      this.$store.dispatch('loadCatchData')
    },
  }
</script>
