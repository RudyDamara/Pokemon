<template>
  <div>
    <v-list-item 
    v-for="(it, index) in item" :key="index"  three-line>
      <v-list-item-content index>
        <div class="text-overline mb-4">
        </div>
        <v-list-item-title 
        @click="clickDetail(it)" class="text-h5 mb-1">
          {{it.name}}
        </v-list-item-title>
      </v-list-item-content>
      <v-list-item-avatar 
        tile
        size="150"> 
        <v-img
        @click="clickDetail(it)"
        :src="it.gambar"></v-img></v-list-item-avatar>
    </v-list-item>
  </div>
</template>

<script>
export default {
    name:"Card",
    props:{
      item: {
        type: [],
        required: false
      },
      clickInformation: {
        type: Function,
        required: false,
    },
    },
      
    data() {
      return {
        itemData:[],
        }
    },
    watch:{
      item(){
        this.itemData = []
        this.itemData = this.item
        console.log(this.item)
      }
    },  
    methods: {
    clickDetail(item) {
      this.$router.push({
        name: "Detail",
        params: {
          data: item,
        },
      });
    },
  },
}
</script>
