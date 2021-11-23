import axios from 'axios'

const header = { 'Access-Control-Allow-Origin': '*' }

export const state = () => ({
  // 検索結果表示関連
  masterMovieList: []
})

export const getters = {
  getMasterMovieList (state) {
    return state.masterMovieList
  }
}

export const mutations = {
  setMasterMovieList (state, movies) {
    state.masterMovieList = movies
  }
}

export const actions = {
  async downloadMasterMovieList (context) {
    const url = `${process.env.apiBaseUrl}/api/v1/masterMovie`
    await axios
      .get(url)
      .then((response) => {
        context.commit('setMasterMovieList', response.data)
      }).catch((error) => {
        throw new Error(error)
      })
  }
}