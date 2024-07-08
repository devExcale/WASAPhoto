<script>
import {axiosConf} from "@/utils/global";
import {User} from "@/utils/entities";

export default {
	data() {
		return {
			userUUID: this.$route.params.user_uuid,
			user: new User(),
			loading: true,
		}
	},
	methods: {

		async load() {

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
		this.load()
	},

}
</script>

<template>
	<div class="container-fluid" v-if="!loading">

		<div class="row">

			<img :src="user.pictureUrl" class="img-thumbnail" alt="...">

			<div class="row">
				<span>{{ user.displayName }}</span>
				<span>{{ user.username }}</span>
			</div>

		</div>

		<ul class="row list-group list-group-horizontal">
			<li class="list-group-item row">
				<p># posts</p>
				<p>{{ user.numPosts }}</p>
			</li>
			<li class="list-group-item row">
				<p># followers</p>
				<p>{{ user.numFollowers }}</p>
			</li>
			<li class="list-group-item row">
				<p># following</p>
				<p>{{ user.numFollowing }}</p>
			</li>
		</ul>

	</div>
	<div class="container-fluid" v-else>
		<h2>Loading...</h2>
	</div>
</template>

<style>
</style>
