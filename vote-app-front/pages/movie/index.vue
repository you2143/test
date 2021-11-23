<template>
  <div>
    <v-btn @click="getdate"></v-btn>
  <v-container>
    <v-row v-for="(data) in datas" :key="data.ID" justify="left" align="left">
      <v-col cols="3" @click="dispData(data)">
        {{data.ID}}
      </v-col>
      <v-col cols="6">
        {{data.title}}
      </v-col>
      <v-spacer></v-spacer>
    </v-row>
  </v-container>
  </div>
</template>

<script>
export default {
  data () {
    return {
      datas: []
    }
  },
  methods: {
    dispData (data) {
      console.log('test')
    },
    getdate () {
      try {
        this.$store.dispatch('movie/downloadMovies')
      } catch (error) {
        throw new Error(error)
      }
      this.datas = this.$store.getters['movie/getMovies']
      console.log('test')
    }
  }
}
</script>
