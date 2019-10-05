<template>
  <div>
    <section class="hero is-warning is-bold">
      <div class="hero-body">
        <div class="container has-text-centered">
          <h1 class="title">
            Upcoming Movies
          </h1>
          <h2 class="subtitle">
            Making you offers you can't refuse
          </h2>
        </div>
      </div>
    </section>

    <div class="box cta small-padding">
      <div class="columns">
        <div class="column"></div>
        <div class="column"></div>
        <div class="column">
          <div class="field has-addons">
            <div class="control">
              <input class="input" type="text" placeholder="Search by name" v-model="typed">
            </div>
            <div class="control" @click="search()">
              <a class="button is-dark">
                Search
              </a>
            </div>
          </div>
        </div>
        <div class="column"></div>
        <div class="column"></div>
      </div>
    </div>

    <movie-list :name="name"></movie-list>

  </div>
</template>

<script>
import MovieList from '@/components/MovieList'

export default {
  name: 'App',
  data () {
    return {
      name: '',
      typed: ''
    }
  },
  computed: {
  },
  methods: {
    search () {
      this.name = this.typed
    }
  },
  components: {
    MovieList
  }
}
</script>

<style lang="sass">
  @import "bulma";

  body {
    background-color: #fff;
  }

  .small-padding {
    padding: .75rem;
  }
</style>
