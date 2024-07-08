<script>
export default {
	props: {
		photoUrl: String,
		username: String,
		caption: String,
		postId: String,
		created: String,
	},
}
</script>

<template>
	<div class="card p-0 m-0" :key="postId">

		<img :src="photoUrl" class="card-img-top" alt="...">
		<div class="card-body">

			<h5 class="card-title">{{ username }}</h5>
			<h6 class="card-subtitle mb-2 text-body-secondary">{{ created }}</h6>
			<p class="card-text">{{ caption }}</p>

			<nav class="nav nav-pills">
				<!-- TODO: change material symbols to feather sprite -->
				<button class="nav-link material-symbols-rounded" @click.prevent="">favorite</button>
				<span class="">5</span>
				<button class="nav-link material-symbols-rounded" @click.prevent="">mode_comment</button>
				<span class="">3</span>
			</nav>

		</div>

	</div>
</template>

<style scoped>

</style>
