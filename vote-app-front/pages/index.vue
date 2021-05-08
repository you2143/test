<template>
  <v-container>
    <v-row v-for="(data, index) in datas" :key="data.num" justify="left" align="left">
      <v-col cols="3" @click="dispData(data)">
        {{index + 1}}
      </v-col>
      <v-col v-show="!data.invisi" cols="6">
        {{data.title}}
      </v-col>
      <v-col v-show="!data.invisi" cols="3">
        {{data.score}}pt
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  data () {
    return {
      datas: [
        {title: "霧の山手線", score: 14, invisi: true},
        {title: "Just out L.A.", score: 10, invisi: true},
        {title: "蹴りたい背中", score: 9, invisi: true},
        {title: "青春おとぼけ白書", score: 8, invisi: true},
        {title: "SKY sea or blue", score: 5, invisi: true}
      ]
    }
  },
  methods: {
    dispData (data) {
      data.invisi = !data.invisi
    }
  }
}
</script>
