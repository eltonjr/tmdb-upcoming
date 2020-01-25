import router from '@/shared/router'

const conf = {
  page: 1,
  EOF: false
}

export default {

  movies: [],

  pending: 0,

  reload () {
    this.clear()
    return this.nextPage()
  },

  clear () {
    conf.page = 1
    conf.EOF = false

    this.movies = []
  },

  nextPage (name) {
    if (conf.EOF || this.pending) {
      return Promise.resolve([])
    }

    this.pending++
    return router.get(router.paths.movies, { name, page: conf.page }).then(res => {
      this.pending--

      if (!res.data) {
        res.data = { movies: [] }
      }

      conf.page++
      conf.EOF = !res.data.movies.length
      this.movies = this.movies.concat(res.data.movies)

      return res.data.movies
    }).catch(failure => {
      this.pending--
      return []
    })
  },

  details (id) {
    this.pending++
    return router.get(`${router.paths.movies}/${id}`).then(res => {
      this.pending--

      if (!res.data) {
        res.data = {}
      }

      return res.data
    }).catch(failure => {
      this.pending--
      return {}
    })
  }
}
