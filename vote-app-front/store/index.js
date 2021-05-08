import axios from 'axios'

const header = { 'Access-Control-Allow-Origin': '*' }

export const state = () => ({
  // 検索結果表示関連
  salonList: [],
  salonCount: 0,
  listPriority: 'created_at__desc',
  salonListPriority: 'created_at',
  listSort: 1,
  currentPage: 1,

  prefecturesMaster: []
})

export const getters = {
  getPrefecturesMaster (state) {
    return state.prefecturesMaster
  },
  getInputFreeWord (state) {
    return state.inputFreeWord
  }
}

export const mutations = {
  setPrefecturesMaster (state, prefectures) {
    state.prefecturesMaster = prefectures
  },

  setSelectAboutCommitmentSunburnSupport (state, isSelect) {
    state.selectAboutCommitment.isSunburnSupport = isSelect
  }
}

export const actions = {
  async downloadCountOfRegisterSalon (context, forSex) {
    forSex = conversionWomen(forSex)
    const url = `${process.env.apiBaseUrl}/${forSex}/api/v1/stores_total`
    await axios
      .get(url)
      .then((response) => {
        context.commit('setCountOfRegisterSalon', response.data.Total)
      }).catch((error) => {
        throw new Error(error)
      })
  },

  async downloadAllSalonsReviews (context, { limit, page, forSex }) {
    forSex = conversionWomen(forSex)
    await axios.get(`${process.env.apiBaseUrl}/${forSex}/api/v1/reviews`, {
      header,
      params: {
        limit,
        page,
        review_status: 'approved'
      }
    }
    ).then((response) => {
      context.commit('setSalonsReviews', response.data.Reviews)
    }).catch((error) => {
      throw new Error(error)
    }).finally(() => {
    })
  },

  async downloadSalonReviews (context, { salonId, limit, page, forSex }) {
    forSex = conversionWomen(forSex)
    await axios.get(`${process.env.apiBaseUrl}/${forSex}/api/v1/stores/${salonId}/reviews`, {
      header,
      params: {
        store_id: salonId,
        limit,
        page,
        review_status: 'approved'
      }
    }
    ).then((response) => {
      context.commit('setSalonReviews', response.data.Reviews)
      context.commit('setSalonReviewCounts', response.data.total)
    }).catch((error) => {
      throw new Error(error)
    }).finally(() => {
    })
  },
}