<template>
  <div>
    <!-- <Card :item= /> -->
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
        @click="Release(items)"
      >
      Release
      </v-btn>

      <v-btn
      class="ma-2"
      color="primary "
      dark
      @click="Rename(items)"
      >
      Rename
      </v-btn>
      
    <v-list-item  three-line>
      <v-list-item-content >
        
        <div class="text-overline mb-4">
        </div>
        <v-list-item-title class="text-h5 mb-1">
          {{items.name}}
        </v-list-item-title>
        
        <b>Moves</b>
        <ul v-for="(it, index) in items.detail.moves" :key="index" >
          <li index>
            {{it.move.name}}
          </li>
        </ul>
        <br>
        <br>
        <b>Types</b>
        <ul v-for="(it, index) in items.detail.types" :key="index" >
          <li index>
            {{it.type.name}}
          </li>
        </ul>
      </v-list-item-content>

      <v-list-item-avatar 
        tile
        size="150"> 
        <v-img
        :src="items.gambar"
      ></v-img> </v-list-item-avatar>
      
    </v-list-item>
  </v-card>

    </div>
</template>

<script>

  export default {
    name: 'MyPokemonDetail',

    components: {
    },

    data() {
      return {
        items:this.$route.params.data,
        }
    },
    computed:{
    },
    methods: {
    back() {
      this.$router.push({
        name: "MyPokemon",
      });
    },
    
    Rename(item) {
      this.$router.push({
        name: "RenamePokemon",
        params: {
          data: item,
        },
      });
    },
    Release(item){
      console.log(item)

        var payload = {
          id: item.id,
        }
        this.$store.dispatch('realase', payload).then(() => {
            if (this.$store.getters.release == true ) {
              this.$router.push({
                name: "MyPokemon"
              });
            } 
        })
    }
  },
    watch:{
    },  
    mounted() {
    },
  }
</script>
