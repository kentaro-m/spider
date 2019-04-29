const axios = require('axios')

module.exports = class ArticlesClient {
  constructor(options) {
    this.options = options
  }

  async post(data) {
    const options = {
      url: '/articles',
      method: 'post',
      baseURL: this.options.baseUrl,
      data: data
    }

    return axios(options)
  }
}