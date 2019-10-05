<template>
  <div class="card is-shady">
    <div class="card-image">
      <figure class="image is-2by3">
        <img src="https://source.unsplash.com/RWnpyGtY1aU" alt="Placeholder image" class="modal-button" data-target="modal-image2">
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
  }
}
</script>

<style lang="sass">
</style>
