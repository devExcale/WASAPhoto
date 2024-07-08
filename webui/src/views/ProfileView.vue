<script>
import {axiosConf, global} from "@/utils/global";
import {User} from "@/utils/entities";
import FeedView from "@/views/FeedView.vue";

export default {
	components: {FeedView},
	data() {
		return {
			userUUID: this.$route.params.user_uuid,
			user: new User(),
			loading: true,
			global: global,
		}
	},
	methods: {

		async loadProfile() {

			this.loading = true;

			try {

				let response = await this.$axios.get(`/users/${this.userUUID}`, axiosConf.value);
				this.user = User.fromResponse(response.data);

				console.log(this.user)

			} catch (e) {

				console.log(e)

				if (e.response && e.response.status === 403) {
					alert('Banned') // TODO
				} else if (e.response && e.response.status === 404) {
					alert('Not Found') // TODO
				} else if (e.response && e.response.status === 500) {
					alert('Internal Server Error')
				} else {
					alert('Generic error.');
				}

			}

			this.loading = false;
		}

	},
	mounted() {
		this.loadProfile()
	},

}
</script>

<template>
	<div class="container-fluid" v-if="!loading">

		<hr>

		<div class="row d-flex">

			<img :src="user.pictureUrl" class="img-thumbnail col text-center" alt="User picture here...">

			<div class="row flex-column col">
				<span class="fw-bold">{{ user.displayName }}</span>
				<span class="fst-italic">{{ user.username }}</span>
			</div>

		</div>

		<ul class="row list-group list-group-horizontal mt-3">
			<li class="list-group-item col text-center">
				<p class="fw-bold h5">posts</p>
				<span class="h5">{{ user.numPosts }}</span>
			</li>
			<li class="list-group-item col text-center">
				<p class="fw-bold h5">followers</p>
				<span class="h5">{{ user.numFollowers }}</span>
			</li>
			<li class="list-group-item col text-center">
				<p class="fw-bold h5">following</p>
				<span class="h5">{{ user.numFollowing }}</span>
			</li>
		</ul>

		<div class="row mt-3" v-if="true"> <!-- TODO same user toolbar-->
			<nav class="nav nav-pills nav-fill justify-content-around">
				<button class="nav-link active m-1" href="/home">follow</button>
				<button class="nav-link active m-1" href="#/">block</button>
			</nav>
		</div>

		<hr>

		<FeedView :userUUID="user.uuid"/>

	</div>
	<div class="container-fluid" v-else>
		<h2>Loading...</h2>
	</div>
</template>

<style>
</style>
