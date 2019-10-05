<template>
  <div>
    <section class="container">
      <div class="columns features" v-for="(row, i) in moviesInFour" :key="i">
        <div class="column is-one-quarter" v-for="(movie, j) in row" :key="j">
          <movie-card :movie="movie" @click.native="select(movie)"></movie-card>
        </div>
      </div>
    </section>

    <movie-details :movie="selected" :showModal="showModal" @close="unselect()"></movie-details>
  </div>
</template>

<script>
import MovieCard from '@/components/MovieCard'
import MovieDetails from '@/components/MovieDetails'

export default {
  name: 'MovieList',
  data () {
    return {
      showModal: false,
      selected: {},
      /* eslint-disable */
      movies: [{"id":522938,"name":"Rambo: Last Blood","poster":"/kTQ3J8oTTKofAVLYnds2cHUz9KO.jpg","genre":"Action"},{"id":540901,"name":"Hustlers","poster":"/zBhv8rsLOfpFW2M5b6wW78Uoojs.jpg","genre":"Comedy"},{"id":453405,"name":"Gemini Man","poster":"/uTALxjQU8e1lhmNjP9nnJ3t2pRU.jpg","genre":"Action"},{"id":431580,"name":"Abominable","poster":"/iK0Q7VWxHsXU1uMzpkf3VjAd6yE.jpg","genre":"Adventure"},{"id":568012,"name":"One Piece: Stampede","poster":"/4E2lyUGLEr3yH4q6kJxPkQUhX7n.jpg","genre":"Action"},{"id":535581,"name":"The Dead Don't Die","poster":"/fgGzTEoNxptCRtEOpOPvIEdlxAq.jpg","genre":"Comedy"},{"id":496243,"name":"Parasite","poster":"/7IiTTgloJzvGI1TAYymCfbfl3vT.jpg","genre":"Comedy"},{"id":567609,"name":"Ready or Not","poster":"/vOl6shtL0wknjaIs6JdKCpcHvg8.jpg","genre":"Comedy"},{"id":504949,"name":"The King","poster":"/8u0QBGUbZcBW59VEAdmeFl9g98N.jpg","genre":"Drama"},{"id":527641,"name":"Five Feet Apart","poster":"/kreTuJBkUjVWePRfhHZuYfhNE1T.jpg","genre":"Drama"},{"id":338967,"name":"Zombieland: Double Tap","poster":"/pIcV8XXIIvJCbtPoxF9qHMKdRr2.jpg","genre":"Action"},{"id":515001,"name":"Jojo Rabbit","poster":"/zIXJoYXc8ezMXc7DEaYii4XwIlS.jpg","genre":"Comedy"},{"id":491283,"name":"Judy","poster":"/5hIORX639YfDlhXfWJTac6bqWlo.jpg","genre":"Drama"},{"id":523077,"name":"Running with the Devil","poster":"/qJ2H94PpXxZgiGGUNkyLSKZzm8u.jpg","genre":"Crime"},{"id":523172,"name":"Late Night","poster":"/xxwb5KBqS0SHsUcyLgWMk3vuVxY.jpg","genre":"Comedy"},{"id":489064,"name":"3 from Hell","poster":"/mxR1tY12XtiMmi3hAjK3f8BRpee.jpg","genre":"Horror"},{"id":420809,"name":"Maleficent: Mistress of Evil","poster":"/vloNTScJ3w7jwNwtNGoG8DbTThv.jpg","genre":"Adventure"},{"id":290859,"name":"Terminator: Dark Fate","poster":"/vqzNJRH4YyquRiWxCCOH0aXggHI.jpg","genre":"Action"},{"id":611788,"name":"The Huntress:  Rune of the Dead","poster":"/qZdBQiU4cEqOrCPyEoBxxM7INsi.jpg","genre":"Action"},{"id":454640,"name":"The Angry Birds Movie 2","poster":"/ebe8hJRCwdflNQbUjRrfmqtUiNi.jpg","genre":"Action"}]
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
  methods: {
    select (movie) {
      this.showModal = true;
      this.selected = movie;
    },
    unselect () {
      this.showModal = false;
    }
  },
  components: {
    MovieCard,
    MovieDetails
  }
}
</script>

<style lang="sass">
</style>
