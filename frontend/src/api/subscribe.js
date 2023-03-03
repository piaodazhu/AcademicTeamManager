import request from '../axios/index'

export function subscribePay(param) {
    return request({
		url: '/subscribe/pay',
		method: 'post',
		data: param,
	})
}

export function getSubscribeInfo(param) {
    return request({
		url: '/subscribe/info',
		method: 'get',
		params: param,
	})
}

