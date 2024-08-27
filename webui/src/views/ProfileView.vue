<script>
import {axiosConf, loadPicture, global} from "@/utils/global";
import {User} from "@/utils/entities";
import FeedView from "@/views/FeedView.vue";

export default {
	components: {FeedView},
	data() {
		return {
			userUUID: this.$route.params.user_uuid,
			user: new User(),
			newUsername: '',
			newDisplayName: '',
			profilePictureSrc: global.loadingGifSrc,
			loading: true,
			global: global,
		}
	},
	inject: ['logout'],
	methods: {

		async loadProfile() {

			this.loading = true;

			try {

				let response = await this.$axios.get(`/users/${this.userUUID}`, axiosConf.value);
				this.user = User.fromResponse(response.data);

				console.log(this.user)

			} catch (e) {

				console.log(e)

				if (e.response && e.response.status === 401) {
					alert('You are not logged in. Try refreshing the page.');
				} else if (e.response && e.response.status === 403) {
					alert('Could not load profile. You are restricted by the user.');
				} else if (e.response && e.response.status === 404) {
					alert('Could not load profile. The user doesn\'t exist.');
				} else if (e.response && e.response.status === 500) {
					alert('Internal Server Error')
				} else {
					alert('Could not load profile. Please try again later. (?)');
				}

			}

			await this.reloadProfilePicture();

			this.loading = false;
		},

		async reloadProfilePicture(src) {

			this.loading = true;
			this.profilePictureSrc = global.loadingGifSrc;

			if (src)
				this.profilePictureSrc = await loadPicture(src);
			else if (this.user.pictureUrl)
				this.profilePictureSrc = await loadPicture(this.user.pictureUrl);

			this.loading = false;
		},

		async changeUsername() {

			this.loading = true;

			try {

				let data = {
					username: this.newUsername,
					displayName: this.newDisplayName,
				};
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

		async deleteProfile() {

			try {

				await this.$axios.delete(`/me`, axiosConf.value);

				await this.logout();

			} catch (e) {

				console.log(e)

				if (e.response && e.response.status === 401) {
					alert('You are not logged in. Try refreshing the page.');
				} else if (e.response && e.response.status === 500) {
					alert('Internal Server Error')
				} else {
					alert('Could not delete profile. Please try again later. (?)');
				}

			}
		},

		async followAction() {

			if (this.loading)
				return;

			this.loading = true;

			if (this.user.loggedUserFollowed)
				await this.unfollowUser();
			else
				await this.followUser();

			await this.loadProfile();

		},

		async followUser() {

			try {

				await this.$axios.put(`/me/followed_users/${this.userUUID}`, {}, axiosConf.value);
				this.user.loggedUserFollowed = true;

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
				this.user.loggedUserFollowed = false;

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

			if (this.user.loggedUserRestricted)
				await this.unrestrictUser();
			else
				await this.restrictUser();

			await this.loadProfile();

		},

		async restrictUser() {

			try {

				await this.$axios.put(`/me/banned_users/${this.userUUID}`, {}, axiosConf.value);
				this.user.loggedUserRestricted = true;

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
				this.user.loggedUserRestricted = false;

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
	provide() {
		return {

			onChangeProfilePicture: this.reloadProfilePicture,

		}
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
			return (this.user.loggedUserFollowed) ? 'unfollow' : 'follow';
		},

		restrictBtnStr() {
			return (this.user.loggedUserRestricted) ? 'unrestrict' : 'restrict';
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

			<img v-if="this.profilePictureSrc"
				 :src="this.profilePictureSrc"
				 @set-profile-picture="src => reloadProfilePicture(src)"
				 class="img-thumbnail col text-center" alt="User picture here...">

			<div v-else class="col text-center">
				<p>Select an image from your profile to show it here!</p>
			</div>

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

		<!-- Actions on logged user -->
		<div class="row" v-else>

			<nav class="nav nav-pills nav-fill justify-content-around mt-3">
				<button class="btn btn-primary nav-link active m-1" type="button" :class="disableOnLoad"
						data-bs-toggle="modal" data-bs-target="#userSettingsModal">
					Options
				</button>
				<RouterLink to="/follows" class="nav-link active m-1">
					Followed Users
				</RouterLink>
				<RouterLink to="/restricts" class="nav-link active m-1">
					Restricted Users
				</RouterLink>
			</nav>

			<!--  -->
			<div class="modal fade" id="userSettingsModal" tabindex="-1">
				<div class="modal-dialog">
					<div class="modal-content">
						<div class="modal-header">

							<h1 class="modal-title fs-5" id="exampleModalLabel">Options</h1>
							<button type="button" class="btn-close" data-bs-dismiss="modal"/>

						</div>

						<div class="modal-body">

							<p class="text-center">
								Here you can change your username and your display name! <br>
								Keep in mind: if you leave a field blank it won't be updated.
							</p>

							<div class="input-group mb-3">
								<div class="input-group-text">@</div>
								<input type="text" class="form-control" placeholder="New Username"
									   v-model="newUsername" />
							</div>

							<input type="text" class="form-control mb-3" placeholder="New Display Name"
								   v-model="newDisplayName" />

						</div>
						<div class="modal-footer">

							<button type="button" class="btn btn-danger" data-bs-dismiss="modal" @click="deleteProfile">Delete profile</button>
							<button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Close</button>
							<button type="button" class="btn btn-primary" data-bs-dismiss="modal" @click="changeUsername">Save changes</button>

						</div>
					</div>
				</div>
			</div>

		</div>

		<hr>

		<FeedView :userUUID="user.uuid"/>

	</div>
</template>

<style>
</style>
