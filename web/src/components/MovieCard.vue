<template>
  <div class="card is-shady">
    <div class="card-image">
      <figure class="image is-2by3">
        <img :src="src(movie.poster)" alt="Movie Poster" class="modal-button" data-target="modal-image2">
      </figure>
    </div>
    <div class="card-content">
      <p class="title is-5">{{movie.name}}</p>
    </div>
  </div>
</template>

<script>
export default {
  name: 'MovieCard',
  props: {
    movie: Object
  },
  data () {
    return {
    }
  },
  methods: {
    src (img) {
      return '/server/images' + img
    }
  }
}
</script>

<style lang="sass">
</style>
