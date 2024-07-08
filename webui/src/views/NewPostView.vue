<script>
import {axiosConfMultipart, updateLogin} from "@/utils/global";
import {Post} from "@/utils/entities";

export default {
	name: "NewPostView",
	data() {
		return {
			photo: null,
			caption: '',
			publishedPost: null,
		};
	},
	methods: {
		async publish() {

			let photoFile = this.$refs.photo.files[0];
			let caption = this.caption;

			if (photoFile === undefined) {
				alert('Photo is required');
				return;
			}

			try {

				let response = await this.$axios.put("/me/feed/", {
					file: photoFile,
					body: JSON.stringify({
						caption: caption
					})
				}, axiosConfMultipart.value);
				this.publishedPost = Post.fromResponse(response.data);

			} catch (e) {

				console.log(e);

				if (e.response && e.response.status === 400) {
					alert('Could not publish post. Please try again later.');
				} else if (e.response && e.response.status === 401) {
					alert('You are not logged in. Try refreshing the page.');
				} else if (e.response && e.response.status === 500) {
					alert('Internal Server Error');
				} else {
					alert('Could not publish post. Please try again later.');
				}

			}
		}
	}
}
</script>

<template>
	<div class="container">

		<p class="text-center h3">Publish a new photo.</p>
		<hr>

		<div class="row mb-auto">

			<!-- File input -->
			<div class="mb-3">
				<label for="input-photo" class="form-label">Photo</label>
				<input class="form-control" type="file" id="input-photo" ref="photo">
			</div>

			<!-- Caption input -->
			<div class="mb-3">
				<label for="input-caption" class="form-label">Caption</label>
				<textarea class="form-control" id="input-caption" rows="3" v-model="caption"></textarea>
			</div>

			<!-- Submit button -->
			<div class="mb-3">
				<button class="btn btn-primary ms-auto me-0" @click="publish">Publish</button>
			</div>

			<hr>

			<div class="row" v-if="publishedPost">
				<p>Post published! <RouterLink to="#">Check it out!</RouterLink></p>
			</div>

		</div>

	</div>
</template>

<style scoped>

</style>
