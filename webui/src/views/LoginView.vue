<script>
import {updateLogin} from "@/utils/global.js";

export default {
	data() {
		return {
			username: '',
		};
	},
	methods: {

		validateUsername() {

			if (this.username === '') {
				return 'Username is required';
			}

			if (this.username.length < 2) {
				return 'Username must be at least 2 characters';
			}

			return null;
		},

		async doLogin() {

			let errMsg = this.validateUsername(this.username);
			if (errMsg !== null && errMsg !== '') {
				alert(errMsg);
				return;
			}

			let user = {
				username: this.username,
			};

			try {

				let response = await this.$axios.post("/session", user);
				let token = response.data.token;
				let userUUID = response.data.user_uuid;

				localStorage.setItem('token', token);
				localStorage.setItem('user_uuid', userUUID);

				updateLogin();

				await this.$router.push('/home');

			} catch (e) {

				console.log(e);

				if (e.response && e.response.status === 500) {
					alert('Internal Server Error');
				} else {
					alert('Could not login. Please try again later.');
				}

			}
		}

	},
}
</script>

<template>
	<div class="justify-content-between m-auto">

		<form @submit.prevent="doLogin" class="row">
			<div class="mb-3">
				<label for="username" class="form-label">Username</label>
				<input type="text" class="form-control" id="username" v-model="username">
			</div>
			<button type="submit" class="btn btn-primary">Login</button>
		</form>
	</div>
</template>

<style>
</style>
