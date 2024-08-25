import {computed, reactive} from "vue";
import axios from '../services/axios.js';

const global = reactive({
	token: '',
	userUUID: '',
	loggedIn: false,
	loadingGifSrc: 'loading.webp',
})

function updateLogin() {
	const token = localStorage.getItem("token");
	const userUUID = localStorage.getItem("user_uuid");

	if (token && userUUID) {
		global.token = token;
		global.userUUID = userUUID;
		global.loggedIn = true;
	} else {
		global.token = '';
		global.userUUID = '';
		global.loggedIn = false;
	}

	return global.loggedIn;
}

async function loadPicture(pictureSrc) {

	let outSrc = global.loadingGifSrc;

	try {

		let response = await axios.get(pictureSrc, axiosConfBlob.value);
		let blob = new Blob([response.data]);
		outSrc = URL.createObjectURL(blob);

	} catch (e) {

		console.log(e)

		if (e.response && e.response.status === 401) {
			alert('You are not logged in. Try refreshing the page.');
		} else if (e.response && e.response.status === 403) {
			alert('Could not load picture. You are restricted by the user.');
		} else if (e.response && e.response.status === 404) {
			alert('Could not load picture. The picture doesn\'t exist.');
		} else if (e.response && e.response.status === 500) {
			alert('Internal Server Error')
		} else {
			alert('Could not load picture. Please try again later. (?)');
		}

	}

	return outSrc;
}

const axiosConf = computed(() => {

	if (!global.loggedIn)
		return {};

	return {
		headers: {
			Authorization: `Bearer ${global.token}`
		}
	}
})

const axiosConfMultipart = computed(() => {

	if (!global.loggedIn)
		return {};

	return {
		headers: {
			Authorization: `Bearer ${global.token}`,
			'Content-Type': 'multipart/form-data',
		}
	}
})

const axiosConfBlob = computed(() => {

	if (!global.loggedIn)
		return {};

	return {
		responseType: 'blob',
		headers: {
			Authorization: `Bearer ${global.token}`
		}
	}
})

export {global, updateLogin, loadPicture, axiosConf, axiosConfMultipart, axiosConfBlob}
