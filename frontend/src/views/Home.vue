<template>
  <div>
    <b>Pokemon</b>
    <Card :item=doneData />
    </div>
</template>

<script>
  import Card from '../components/Card'

  export default {
    name: 'Home',

    components: {
      Card,
    },
    computed:{
      doneData () {
        return this.$store.getters.items
      }
    },
    methods: {
    clickDetail(param) {
      this.$router.push({
        name: "Detail",
        params: {
          data: param,
        },
      });
    },
    },
    watch:{
      doneData(){
        this.itemPokemon= this.doneData
        console.log("data",this.doneData)
      }
    },  
    data() {
      return {
        data:null,
        }
    },
    mounted() {
      this.$store.dispatch('loadData') // dispatch loading
    },
  }
</script>
