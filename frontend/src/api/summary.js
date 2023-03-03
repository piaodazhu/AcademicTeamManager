import request from '../axios/index'

export function getSummary(param) {
    return request({
		url: '/summary',
		method: 'get',
		params: param,
	})
}
