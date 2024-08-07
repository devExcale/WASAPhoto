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
			newUsername: '',
			loading: true,
			global: global,
			followStatus: false,
			restrictStatus: false,
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
					alert('Could not load profile. You are restricted by the user.');
				} else if (e.response && e.response.status === 404) {
					alert('Could not load profile. The user doesn\'t exist.');
				} else if (e.response && e.response.status === 500) {
					alert('Internal Server Error')
				} else {
					alert('Could not load profile. Please try again later. (?)');
				}

			}

			this.loading = false;
		},

		async changeUsername() {

			this.loading = true;

			try {

				let data = {username: this.newUsername};
				await this.$axios.patch(`/me/username`, data, axiosConf.value);

				console.log(this.user)

			} catch (e) {

				console.log(e)

				if (e.response && e.response.status === 400) {
					alert('Could not change username. Please try again later.');
				} else if (e.response && e.response.status === 401) {
					alert('You are not logged in. Try refreshing the page.');
				} else if (e.response && e.response.status === 409) {
					alert('The requested username is already taken.');
				} else if (e.response && e.response.status === 500) {
					alert('Internal Server Error')
				} else {
					alert('Could not change username. Please try again later. (?)');
				}

			}

			await this.loadProfile();
		},

		async followAction() {

			if (this.loading)
				return;

			this.loading = true;

			if (this.followStatus)
				await this.unfollowUser();
			else
				await this.followUser();

			await this.loadProfile();

		},

		async followUser() {

			try {

				await this.$axios.put(`/me/followed_users/${this.userUUID}`, {}, axiosConf.value);
				this.followStatus = true;

				console.log(`User ${this.userUUID} followed.`)

			} catch (e) {

				console.log(e)

				if (e.response && e.response.status === 400) {
					alert('Could not follow. Please try again later.');
				} else if (e.response && e.response.status === 401) {
					alert('You are not logged in. Try refreshing the page.');
				} else if (e.response && e.response.status === 403) {
					alert('Could not follow. You are restricted by the user.');
				} else if (e.response && e.response.status === 404) {
					alert('Could not follow. The user doesn\'t exist.');
				} else if (e.response && e.response.status === 500) {
					alert('Internal Server Error')
				} else {
					alert('Could not follow. Please try again later. (?)');
				}

			}

		},

		async unfollowUser() {

			try {

				await this.$axios.delete(`/me/followed_users/${this.userUUID}`, axiosConf.value);
				this.followStatus = false;

				console.log(`User ${this.userUUID} unfollowed.`)

			} catch (e) {

				console.log(e)

				if (e.response && e.response.status === 400) {
					alert('Could not unfollow. Please try again later.');
				} else if (e.response && e.response.status === 401) {
					alert('You are not logged in. Try refreshing the page.');
				} else if (e.response && e.response.status === 404) {
					alert('Could not unfollow. The user doesn\'t exist.');
				} else if (e.response && e.response.status === 500) {
					alert('Internal Server Error')
				} else {
					alert('Could not unfollow. Please try again later. (?)');
				}

			}

		},

		async restrictAction() {

			if (this.loading)
				return;

			this.loading = true;

			if (this.restrictStatus)
				await this.unrestrictUser();
			else
				await this.restrictUser();

			await this.loadProfile();

		},

		async restrictUser() {

			try {

				await this.$axios.put(`/me/banned_users/${this.userUUID}`, {}, axiosConf.value);
				this.restrictStatus = true;

				console.log(`User ${this.userUUID} restricted.`)

			} catch (e) {

				console.log(e)

				if (e.response && e.response.status === 400) {
					alert('Could not restrict. Please try again later.');
				} else if (e.response && e.response.status === 401) {
					alert('You are not logged in. Try refreshing the page.');
				} else if (e.response && e.response.status === 404) {
					alert('Could not restrict. The user doesn\'t exist.');
				} else if (e.response && e.response.status === 500) {
					alert('Internal Server Error')
				} else {
					alert('Could not restrict. Please try again later. (?)');
				}

			}

		},

		async unrestrictUser() {

			try {

				await this.$axios.delete(`/me/banned_users/${this.userUUID}`, axiosConf.value);
				this.restrictStatus = false;

				console.log(`User ${this.userUUID} restricted.`)

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
					alert('Could not restrict. Please try again later. (?)');
				}

			}

		},

	},
	mounted() {
		this.loadProfile()
	},
	computed: {

		isLoggedUser() {
			return this.userUUID === this.global.userUUID;
		},

		disableOnLoad() {
			return {
				'disabled': this.loading
			}
		},

		followBtnStr() {
			return (this.followStatus) ? 'unfollow' : 'follow';
		},

		restrictBtnStr() {
			return (this.restrictStatus) ? 'unrestrict' : 'restrict';
		},

	}

}
</script>

<template>
	<div class="container-fluid" v-if="loading">
		<h2>Loading...</h2>
	</div>
	<div class="container-fluid" v-else>

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

		<div class="row mt-3" v-if="!isLoggedUser">
			<!-- Actions on other users -->
			<nav class="nav nav-pills nav-fill justify-content-around">
				<button :class="disableOnLoad" class="nav-link active m-1" @click="followAction">
					{{ followBtnStr }}
				</button>
				<button :class="disableOnLoad" class="nav-link active m-1" @click="restrictAction">
					{{ restrictBtnStr }}
				</button>
			</nav>
		</div>

		<div class="row" v-else>
			<!-- Actions on logged user -->
			<nav class="nav nav-pills nav-fill justify-content-around mt-3">
				<RouterLink to="/follows" class="nav-link active m-1">
					followed users
				</RouterLink>
				<RouterLink to="/restricts" class="nav-link active m-1">
					restricted users
				</RouterLink>
			</nav>

			<div class="input-group mt-3">
				<input type="text" class="form-control" placeholder="New Username" v-model="newUsername">
				<button class="btn btn-primary" type="button" @click="changeUsername" :class="disableOnLoad">
					Change Username
				</button>
			</div>

		</div>

		<hr>

		<FeedView :userUUID="user.uuid"/>

	</div>
</template>

<style>
</style>
