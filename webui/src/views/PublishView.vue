<script>
import {axiosConfMultipart, global} from "@/utils/global";
import {Post} from "@/utils/entities";
import router from "@/router";

export default {
	name: "PublishView",
	data() {
		return {
			photo: null,
			caption: '',
			publishedPost: null,
		};
	},
	computed: {
		previewURL() {
			let a = this.photo ? URL.createObjectURL(this.photo) : null;
			console.debug('previewURL: ' + a)
			return a;
		}
	},
	methods: {

		getPhotoFile() {
			this.photo = this.$refs.photo.files[0];
		},

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

				await router.push(`/profile/${global.userUUID}`);

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

		<p class="text-center h3">Publish a new photo</p>
		<hr>

		<div class="row mb-auto">

			<!-- File input -->
			<div class="mb-3 d-none">
				<label for="input-photo" class="form-label h6">File</label>
				<input class="form-control" type="file" id="input-photo" ref="photo" @change="getPhotoFile">
			</div>

			<!-- Photo preview -->
			<div class="mb-3" style="max-height: 50vh">
				<img :src="previewURL" id="preview-photo" alt="Photo preview" v-if="previewURL"
						class="border border-3 border-dark rounded-4 d-block m-0 mx-auto"
						style="max-height: 100%; max-width: 100%;"
						@click="$refs.photo.showPicker()"/>
				<div class="text-center w-100 m-0 py-5 border border-1 rounded-4 d-block" v-else
						@click="$refs.photo.showPicker()">
					<span class="h5 fst-italic">Click to select a photo</span>
				</div>
			</div>

			<!-- Tip -->
			<p class="fst-italic m-0 mb-3 text-center" v-if="previewURL">
				Tip: Click the image again to select a different photo.
			</p>

			<!-- Caption input -->
			<div class="mb-3">
				<div class="form-floating">
					<textarea class="form-control" id="input-caption" placeholder="Caption" v-model="caption"
							style="height: 10em;" />
					<label for="input-caption" class="form-label h6">Caption</label>
				</div>
			</div>

			<!-- Submit button -->
			<div class="mb-3">
				<button class="btn btn-primary m-0 w-100" @click="publish">Publish</button>
			</div>

			<hr>

			<div class="row" v-if="previewURL">
			</div>

		</div>

	</div>
</template>

<style scoped>

</style>
