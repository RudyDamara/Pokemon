<template>
  <card-detail :item="items" />
</template>

<script>
  import CardDetail from '../components/CardDetail'

  export default {
    name: 'Detail',
    components: {
      CardDetail,
    },
    methods: {
      // Back(data) {
      // },
    },

    watch:{
      items(){
        console.log(this.items)
      }
    },  
    data() {
      return {
        items:this.$route.params.data,
        }
    },
  }
</script>
