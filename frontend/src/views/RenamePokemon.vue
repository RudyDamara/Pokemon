<template>
  <div>
    <v-card
    class="mx-auto"
    max-width="500"
    outlined
  >
  <!-- {{items}} -->
    <v-text-field
      v-model="name"
      hide-details="auto"
    ></v-text-field>
    <v-btn
        class="ma-2"
        color="primary "
        dark
        @click="Save()"
      >Save
      </v-btn>
    </v-card>
  </div>
</template>

<script>

  export default {
    name: 'RenamePokemon',

    components: {
    },

    data() {
      return {
        items:this.$route.params.data,
        name:this.$route.params.data.name
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
    Save(){

        var payload = {
          id: this.items.id,
          name: this.name,
        }
        console.log("Sace")
        this.$store.dispatch('update', payload).then(() => {
            if (this.$store.getters.update == true ) {
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
