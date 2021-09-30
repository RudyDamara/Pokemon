<template>
  <v-card
    class="mx-auto"
    max-width="500"
    outlined
  >
    <v-btn
        class="ma-2"
        color="primary "
        dark
        @click="back()"
      >
        <v-icon
          dark
          left
        >
          mdi-arrow-left
        </v-icon>Back
      </v-btn>
      
      <v-btn
        class="ma-2"
        color="primary "
        dark
        @click="Catch(item)"
      >
      Catch
      </v-btn>
      
    <v-list-item  three-line>
      <v-list-item-content >
        
        <div class="text-overline mb-4">
        </div>
        <v-list-item-title class="text-h5 mb-1">
          {{item.name}}
        </v-list-item-title>
        
        <b>Moves</b>
        <ul v-for="(it, index) in item.detail.moves" :key="index" >
          <li index>
            {{it.move.name}}
          </li>
        </ul>
        <br>
        <br>
        <b>Types</b>
        <ul v-for="(it, index) in item.detail.types" :key="index" >
          <li index>
            {{it.type.name}}
          </li>
        </ul>
      </v-list-item-content>

      <v-list-item-avatar 
        tile
        size="150"> 
        <v-img
        :src="item.gambar"
      ></v-img> </v-list-item-avatar>
      
    </v-list-item>
  </v-card>
</template>

<script>
export default {
    name:"Card",
    props:{
      item: {
        type: Object,
        required: false
      },
      
    },
    
    methods: {
    back() {
      this.$router.push({
        name: "Home",
      });
    },
    Catch(item){
      console.log(item)

        var payload = {
          name: item.name,
          url: item.url,
        }
        this.$store.dispatch('saveData', payload).then(() => {
            if (this.$store.getters.responseCatchPokemon == true ) {
              this.$router.push({
                name: "MyPokemon"
              });
            } 
        })
    }
  },
  
  
}
</script>
