<script>
import {User} from "@/utils/entities";
import {axiosConf} from "@/utils/global";

export default {
	name: "RestrictedView",
	data: function () {
		return {
			users: [],
			loading: false,
		}
	},
	methods: {

		async loadData() {

			try {

				let response = await this.$axios.get('/me/banned_users/', axiosConf.value);
				this.users = response.data.map(data => User.fromResponse(data));

				console.log(response)
				console.log(this.users)

			} catch (e) {

				console.log(e)

				if (e.response && e.response.status === 401) {
					alert('You are not logged in. Try refreshing the page.');
				} else if (e.response && e.response.status === 500) {
					alert('Internal Server Error');
				} else {
					alert('Could not search. Please try again later.');
				}

			}
		},

		async unrestrictUser(user) {

			if (this.loading)
				return;

			this.loading = true;

			try {

				await this.$axios.delete(`/me/banned_users/${user.uuid}`, axiosConf.value);
				this.users = this.users.filter(u => u.uuid !== user.uuid);

				console.log(`User ${this.userUUID} unrestricted.`)

			} catch (e) {

				console.log(e)

				if (e.response && e.response.status === 400) {
					alert('Could not unrestrict. Please try again later.');
				} else if (e.response && e.response.status === 401) {
					alert('You are not logged in. Try refreshing the page.');
				} else if (e.response && e.response.status === 404) {
					alert('Could not unrestrict. The user doesn\'t exist.');
				} else if (e.response && e.response.status === 500) {
					alert('Internal Server Error')
				} else {
					alert('Could not unrestrict. Please try again later. (?)');
				}

			}

			this.loading = false;

		},

	},
	mounted() {
		this.loadData()
	},
}
</script>

<template>
	<div class="container-fluid">

		<p class="h3 text-center">Restricted Users</p>

		<hr>

		<div v-if="users.length === 0" class="row m-4 mb-3">
			<p class="h4 text-center">You haven't restricted any users.</p>
		</div>


		<div v-for="user in users" :key="user.uuid" class="input-group mb-3 w-100">
			<RouterLink :to="`/profile/${user.uuid}`" class="input-group-text">
				{{ user.displayName }} - {{ user.username }}
			</RouterLink>
			<button :class="{'disabled': this.loading}" class="btn btn-outline-secondary" type="button"
					@click="unrestrictUser(user)">X
			</button>
		</div>

	</div>
</template>

<style>
</style>
