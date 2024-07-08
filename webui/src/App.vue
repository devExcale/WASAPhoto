<script>
import {global, updateLogin} from "@/utils/global.js";

export default {
	data() {
		return {
			global: global,
		};
	},
	methods: {

		async logout() {

			localStorage.removeItem('token');
			localStorage.removeItem('user_uuid');

			updateLogin();

			await this.$router.push('/login');
		}

	},
	mounted() {

		updateLogin()

		if (!this.global.loggedIn)
			this.$router.push('/login');

	}
}
</script>

<template>
	<div class="container-fluid vh-100 bg-dark p-0">
		<div
			class="d-flex flex-column m-auto justify-content-center vh-100 col-md-4 border border-primary bg-white p-0">

			<div class="row text-center bg-primary">
				<div class="col text-light">
					<h1>WASAPhoto</h1>
				</div>
			</div>

			<div class="row p-3 flex-grow-1 d-flex overflow-scroll">
				<div class="col d-flex justify-content-center">
					<RouterView/>
				</div>
			</div>

			<div class="row border-top border-primary" v-if="global.loggedIn">
				<nav class="nav nav-pills nav-fill justify-content-around">
					<!-- TODO: change material symbols to feather sprite -->
					<a class="nav-link rounded-0 active material-symbols-rounded" href="/home">home</a>
					<a class="nav-link rounded-0 material-symbols-rounded" href="#/">search</a>
					<RouterLink to="/newPost"
								class="nav-link rounded-0 material-symbols-rounded">
						add_box
					</RouterLink>
					<RouterLink :to="`/profile/${global.userUUID}`"
								class="nav-link rounded-0 material-symbols-rounded">
						person
					</RouterLink>
					<button class="nav-link rounded-0 material-symbols-rounded" @click.prevent="logout">logout</button>
				</nav>
			</div>

		</div>
	</div>
</template>

<style>
</style>
