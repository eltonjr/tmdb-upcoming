<template>
  <div class="modal" :class="{'is-active': showModal, 'is-clipped': showModal}">
    <div class="modal-background" @click="close()"></div>
    <div class="modal-content">
      <div class="box">
        <div class="columns">
          <div class="column is-two-fifths">
            <figure class="media-left">
              <p class="image is-2by3">
                <img src="https://source.unsplash.com/RWnpyGtY1aU">
              </p>
            </figure>
          </div>
          <div class="column">
            <div class="content">
              <h1 class="title is-size-3">{{movie.name}}</h1>
              <p class="is-size-5 subtitle" v-if="movie.genre">
                Genre: <strong>{{movie.genre}}</strong>
              </p>
              <p class="is-size-5 subtitle" v-if="movie.releaseDate">
                Release Date: <strong>{{movie.releaseDate}}</strong>
              </p>
              <p class="is-size-5 subtitle" v-if="movie.overview">
                {{movie.overview}}
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
    <button class="modal-close is-large" aria-label="close" @click.stop="close()"></button>
  </div>
</template>

<script>
export default {
  name: 'MovieDetails',
  props: {
    showModal: Boolean,
    movie: Object
  },
  data () {
    return {
    }
  },
  methods: {
    close () {
      this.$emit('close')
    }
  }
}
</script>

<style lang="sass">
</style>
