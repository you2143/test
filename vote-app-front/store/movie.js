import axios from 'axios'

export const state = () => ({
  movies: [],
  movie: {},
  movieCount: 0
})

export const getters = {
  getMoviesCount (state) {
    return state.movieCount
  },
  getMovies (state) {
    return state.movies
  },
  getMovie (state) {
    return state.movie
  }
}

export const mutations = {
  setMovies (state, movies) {
    state.movies = movies
  },
  setMovie (state, movie) {
    state.movie = movie
  },
  initMovies (state) {
    state.movies = []
  },
  initMovie (state) {
    state.movie = {}
  },
  setMovieAttribute (state, { field, value }) {
    state.Movie[field] = value
  }
}

export const actions = {
  downloadMovies (context) {
    const url = `http://localhost:28080/api/v1/masterMovie`
    axios.get(url)
      .then((res) => {
        context.commit('setMovies', res.data)
      })
  },
  downloadMovie (context, payload) {
    const url = `${process.env.apiBaseUrl}/api/v1/masterMovie/${payload.id}`

    return axios
      .get(url)
      .then((res) => {
        context.commit('setMovie', res.data)
      })
  },
  deleteMovie (context, payload) {
    const url = `${process.env.apiBaseUrl}/api/v1/masterMovie/${payload.id}`

    return axios
      .delete(url)
      .catch((error) => {
        throw new Error(error)
      })
      .finally(() => {})
  },
  uploadMovies (context, payload) {
    const url = `${process.env.apiBaseUrl}/api/v1/masterMovie/${payload.id}`
    const params = new FormData()
    params.append('title', payload.title)
    return axios
      .put(url, params, {
        headers: {
          'content-type': 'multipart/form-data'
        }
      })
      .then((res) => {
        context.commit('setMovie', res.data)
      })
  },
  createMovies (context, payload) {
    const url = `${process.env.apiBaseUrl}/api/v1/masterMovie`
    const params = new FormData()
    params.append('title', payload.title)
    return axios
      .post(url, params, {
        headers: {
          'content-type': 'multipart/form-data'
        }
      })
      .then((res) => {
        context.commit('setMovie', res.data)
      })
  }
}
