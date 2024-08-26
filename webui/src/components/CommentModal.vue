<script>
import {Post, Comment} from "@/utils/entities";
import {axiosConf} from "@/utils/global";

export default {
	name: "CommentModal",
	data() {
		return {
			comments: [],
			newCommentInput: '',
		}
	},
	props: {
		post: Post,
	},
	computed: {
	},
	methods: {

		async loadComments() {

			try {

				let url = `/users/${this.post.authorUuid}/feed/${this.post.uuid}/comments/`;
				let response = await this.$axios.get(url, axiosConf.value);
				this.comments = response.data.map(Comment.fromResponse);

			} catch (e) {

				console.log(e)

				if (e.response && e.response.status === 400) {
					alert('Could not create comment. Please try again later.');
				} else if (e.response && e.response.status === 401) {
					alert('You are not logged in. Try refreshing the page.');
				} else if (e.response && e.response.status === 403) {
					alert('Could not create comment. You are restricted by the user.');
				} else if (e.response && e.response.status === 404) {
					alert('Could not create comment. The post doesn\'t exist.');
				} else if (e.response && e.response.status === 500) {
					alert('Internal Server Error')
				} else {
					alert('Could not create comment. Please try again later. (?)');
				}

			}
		},

		async commentPost() {

			if (this.newCommentInput.length === 0) {
				alert('Comment cannot be empty.');
				return;
			}

			try {

				let url = `/users/${this.post.authorUuid}/feed/${this.post.uuid}/comments/`;
				let response = await this.$axios.put(url, {"comment": this.newCommentInput}, axiosConf.value);
				this.comments.push(Comment.fromResponse(response.data));
				this.post.numComments++;
				this.newCommentInput = '';

			} catch (e) {

				console.log(e)

				if (e.response && e.response.status === 400) {
					alert('Could not create comment. Please try again later.');
				} else if (e.response && e.response.status === 401) {
					alert('You are not logged in. Try refreshing the page.');
				} else if (e.response && e.response.status === 403) {
					alert('Could not create comment. You are restricted by the user.');
				} else if (e.response && e.response.status === 404) {
					alert('Could not create comment. The post doesn\'t exist.');
				} else if (e.response && e.response.status === 500) {
					alert('Internal Server Error')
				} else {
					alert('Could not create comment. Please try again later. (?)');
				}

			}
		},

	},
	mounted() {
		this.loadComments();
	}
}
</script>

<template>
	<div class="modal fade" :id="'comments-' + post.uuid" tabindex="-1">
		<div class="modal-dialog modal-dialog-centered modal-dialog-scrollable modal-xl">
			<div class="modal-content">

				<div class="modal-header">
					<h1 class="modal-title fs-5">Comments</h1>
					<button type="button" class="btn-close" data-bs-dismiss="modal"></button>
				</div>

				<div class="modal-body">

					<p v-if="this.comments.length === 0" class="text-center h5">
						No comments yet...
					</p>

					<ul v-else>
						<li v-for="comment in this.comments" :key="comment.uuid">
							<span class="fst-italic fw-bold">{{ comment.authorUuid }}</span> {{ comment.comment }}
						</li>
					</ul>

				</div>

				<div class="modal-footer">

					<div class="input-group">
						<input type="text" class="form-control" placeholder="Type your comment here"
							   v-model="newCommentInput" @keydown.enter="commentPost"/>
						<button type="button" class="btn btn-primary" @click="commentPost">Submit</button>
					</div>

				</div>

			</div>
		</div>
	</div>
</template>

<style scoped>

</style>
