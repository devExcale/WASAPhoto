<script>
import {Post, Comment} from "@/utils/entities";
import {axiosConf, global} from "@/utils/global";

export default {
	name: "CommentModal",
	data() {
		return {
			comments: [],
			newCommentInput: '',
			modal: null,
		}
	},
	props: {
		post: Post,
	},
	computed: {
		global() {
			return global
		},
		modalId() {
			return `comments-${this.post.uuid}`;
		},
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

		async deleteComment(uuid) {

			try {

				let url = `/users/${this.post.authorUuid}/feed/${this.post.uuid}/comments/${uuid}`;
				await this.$axios.delete(url, axiosConf.value);

				this.comments = this.comments.filter(c => c.uuid !== uuid);
				this.post.numComments--;

			} catch (e) {

				console.warn(e)

				if (e.response && e.response.status === 400) {
					alert('Could not delete comment. Please try again later.');
				} else if (e.response && e.response.status === 401) {
					alert('You are not logged in. Try refreshing the page.');
				} else if (e.response && e.response.status === 403) {
					alert('Could not delete comment. You are restricted by the user.');
				} else if (e.response && e.response.status === 404) {
					alert('Could not delete comment. The post doesn\'t exist.');
				} else if (e.response && e.response.status === 500) {
					alert('Internal Server Error')
				} else {
					alert('Could not delete comment. Please try again later. (?)');
				}

			}

		}

	},
	mounted() {
		this.loadComments();
		this.modal = new bootstrap.Modal('#' + this.modalId, {keyboard: false});
	}
}
</script>

<template>
	<div class="modal fade" :id="modalId" tabindex="-1">
		<div class="modal-dialog modal-dialog-centered modal-dialog-scrollable modal-xl">
			<div class="modal-content">

				<div class="modal-header">
					<h1 class="modal-title fs-5">Comments</h1>
					<button type="button" class="btn-close" data-bs-dismiss="modal"></button>
				</div>

				<div class="modal-body">

					<div class="container" v-if="this.comments.length === 0">
						<hr>
						<p class="text-center h5">
							No comments yet...
						</p>
						<hr>
					</div>

					<div class="container p-0 text-start legend" v-else>
						<div class="row p-2" v-for="comment in this.comments" :key="comment.uuid">
								<p class="col my-0 mx-1 d-inline">
									<RouterLink :to="`/profile/${comment.authorUuid}`" class="fst-italic fw-bold"
												@click.prevent="modal.hide()" ref="routerLinkTag">
										{{ comment.authorName }}
									</RouterLink>
									&gt;
									<span>{{ comment.comment }}</span>
								</p>
								<button type="button" class="btn-close btn-sm mx-3 my-0"
										@click="deleteComment(comment.uuid)"
										v-if="comment.authorUuid === global.userUUID"></button>
						</div>
					</div>

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
.legend .row:nth-of-type(odd) {
	background-color: azure;
}

.legend .row:nth-of-type(even) {
	background: lavender;
}
</style>
