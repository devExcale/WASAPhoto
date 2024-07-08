<script>
import {User} from "@/utils/entities";
import {axiosConf} from "@/utils/global";

export default {
	name: "SearchView",
	data: function () {
		return {
			query: '',
			users: [],
		}
	},
	props: {
		userUUID: String,
	},
	methods: {

		async doSearch() {

			if (this.query.length < 3) {
				this.users = [];
				return;
			}

			try {

				let response = await this.$axios.get(`/users/?username=${this.query}`, axiosConf.value);
				this.users = response.data.map(data => User.fromResponse(data));

				console.log(response)
				console.log(this.users)

			} catch (e) {

				console.log(e)

				if (e.response && e.response.status === 400) {
					alert('Could not search. Please try again later.');
				} else if (e.response && e.response.status === 401) {
					alert('You are not logged in. Try refreshing the page.');
				} else if (e.response && e.response.status === 500) {
					alert('Internal Server Error');
				} else {
					alert('Could not search. Please try again later.');
				}

			}
		},

	},
	mounted() {
		this.doSearch()
	},
}
</script>

<template>
	<div class="container-fluid">

		<input class="form-control" type="text" placeholder="Username here" v-model="query" @input="doSearch">

		<hr>

		<div v-if="users.length === 0" class="row m-4 mb-3">
			<p class="h4 text-center">Use the search bar to search a user...</p>
		</div>

		<div v-for="user in users" :key="user.uuid" class="row m-4 mb-3">

			<RouterLink :to="`/profile/${user.uuid}`" class="">
				{{ user.displayName }} - {{ user.username }}
			</RouterLink>

			<hr>

		</div>

	</div>
</template>

<style>
</style>
