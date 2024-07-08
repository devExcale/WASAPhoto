<script>
import UserPost from "@/components/UserPost.vue";

export default {
	components: {UserPost},
	data: function () {
		return {
			errormsg: null,
			loading: false,
			some_data: null,
		}
	},
	methods: {
		async refresh() {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get("/");
				this.some_data = response.data;
			} catch (e) {
				// this.errormsg = e.toString();
			}
			this.loading = false;
		},
	},
	mounted() {
		this.refresh()
	}
}
</script>

<template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3">
			<UserPost username="Cr4zy5ky_U"
					  caption="Currently rewatching this"
					  photo-url="https://imgsrv.crunchyroll.com/cdn-cgi/image/fit=contain,format=auto,quality=85,width=1200,height=675/catalog/crunchyroll/2d2c3381c9aabc3acb9551199f383ca3.jpe"
					  post-id="1"
					  created="2021-10-01 12:00:00"
			/>
		</div>
	</div>
</template>

<style>
</style>
