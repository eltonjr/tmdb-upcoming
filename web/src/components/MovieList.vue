<template>
  <div>
    <section class="container">
      <div class="columns features" v-for="(row, i) in moviesInFour" :key="i">
        <div class="column is-one-quarter" v-for="(movie, j) in row" :key="j">
          <movie-card :movie="movie" @click.native="select(movie)"></movie-card>
        </div>
      </div>
      <infinite-loading @infinite="loadMovies">
        <span slot="no-results"></span>
        <span slot="no-more"></span>
      </infinite-loading>
    </section>

    <movie-details :movie="selected" :showModal="showModal" @close="unselect()"></movie-details>
  </div>
</template>

<script>
import InfiniteLoading from 'vue-infinite-loading'

import MovieCard from '@/components/MovieCard'
import MovieDetails from '@/components/MovieDetails'

import service from '@/components/movie'

export default {
  name: 'MovieList',
  props: {
    name: String
  },
  data () {
    return {
      showModal: false,
      pending: 0,
      selected: {},
      movies: []
    }
  },
  computed: {
    moviesInFour () {
      const arr = []
      let i = 0
      while (this.movies && i < this.movies.length) {
        if (i % 4 === 0) arr.push([])
        arr[arr.length - 1].push(this.movies[i++])
      }
      return arr
    }
  },
  created () {
    this.loadMovies()
  },
  methods: {
    select (movie) {
      service.details(movie.id).then(data => {
        this.showModal = true
        this.selected = data
      }).catch(() => {
        console.log('TODO: better error handling')
      })
    },
    unselect () {
      this.showModal = false
    },
    loadMovies ($state) {
      this.pending++
      service.nextPage(this.name).then(data => {
        this.pending--
        this.movies = service.movies
        if ($state) {
          $state.loaded()
          if (!data.length) {
            $state.complete()
          }
        }
      }).catch(() => {
        this.pending--
        if ($state) {
          $state.complete()
        }
      })
    }
  },
  watch: {
    name () {
      service.clear()
      this.loadMovies()
    }
  },
  components: {
    MovieCard,
    MovieDetails,
    InfiniteLoading
  }
}
</script>

<style lang="sass">
</style>
