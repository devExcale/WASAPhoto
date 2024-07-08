import {computed, reactive} from "vue";

const global = reactive({
	token: '',
	userUUID: '',
	loggedIn: false,
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

const axiosConf = computed(() => {

	if (!global.loggedIn)
		return {};

	return {
		headers: {
			Authorization: `Bearer ${global.token}`
		}
	}
})

export {global, updateLogin, axiosConf}
