import request from '../axios/index'

export function userLogin(param) {
    return request({
		url: '/user/login',
		method: 'post',
		data: param,
	})
}

export function userRegister(param) {
    return request({
		url: '/user/register',
		method: 'post',
		data: param,
	})
}

export function getVerifyCode(param) {
    return request({
		url: '/user/verifycode',
		method: 'get',
		params: param,
	})
}

export function userForgotPass(param) {
    return request({
		url: '/user/forgetpass',
		method: 'post',
		data: param,
	})
}

export function userDelete(param) {
    return request({
		url: '/user/delete',
		method: 'post',
		data: param,
	})
}

export function getUserInfo(param) {
    return request({
		url: '/user/info',
		method: 'get',
		params: param,
	})
}