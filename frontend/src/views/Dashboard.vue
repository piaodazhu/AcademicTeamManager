<template>
    <div :style="{ padding: '20px 20px 12px 20px' }">
        <a-row :gutter="16">
            <a-col :span="8">
                <a-card class="card">
                    <a-statistic :value="data.Students" style="margin-right: 50px">
                        <template #title>
                            <span>全部学生</span>
                            <a-tooltip placement="right">
                                <template #title>
                                    <span>学生数量，单位（人）</span>
                                </template>
                                <question-circle-two-tone style="margin-left: 5px" />
                            </a-tooltip>
                        </template>
                    </a-statistic>
                </a-card>
            </a-col>
            <a-col :span="8">
                <a-card class="card">
                    <a-statistic :value="data.projects" style="margin-right: 50px">
                        <template #title>
                            <span>全部项目</span>
                            <a-tooltip placement="right">
                                <template #title>
                                    <span>项目数量，单位（项）</span>
                                </template>
                                <question-circle-two-tone style="margin-left: 5px" />
                            </a-tooltip>
                        </template>
                    </a-statistic>
                </a-card>
            </a-col>
            <a-col :span="8">
                <a-card class="card">
                    <a-statistic :value="data.outputs" style="margin-right: 50px">
                        <template #title>
                            <span>产出数量</span>
                            <a-tooltip placement="right">
                                <template #title>
                                    <span>项目相关学术产出</span>
                                </template>
                                <question-circle-two-tone style="margin-left: 5px" />
                            </a-tooltip>
                        </template>
                    </a-statistic>
                </a-card>
            </a-col>
        </a-row>
        <a-row :gutter="16">
            <a-col :span="8">
                <a-card class="card" style="height: 40vh;margin-top: 20px;">
                    <div style="color: #606266;font-size: 16px;font-weight: 600;">
                        <span>学生团队</span>
                        <a-tooltip placement="right">
                            <template #title>
                                <span>学生学位，单位（个）</span>
                            </template>
                            <question-circle-two-tone style="margin-left: 5px" />
                        </a-tooltip>
                    </div>
                    <div id="Student" style="width: 100%; height: 360px;"></div>
                </a-card>
            </a-col>
            <a-col :span="8">
                <a-card class="card" style="height: 40vh;margin-top: 20px;">
                    <div style="color: #606266;font-size: 16px;font-weight: 600;">
                        <span>项目成果数量</span>
                        <a-tooltip placement="right">
                            <template #title>
                                <span>项目成果，单位（项）</span>
                            </template>
                            <question-circle-two-tone style="margin-left: 5px" />
                        </a-tooltip>
                    </div>
                    <div id="Project" style="width: 100%; height: 360px;"></div>
                </a-card>
            </a-col>
            <a-col :span="8">
                <a-card class="card" style="height: 40vh;margin-top: 20px;">
                    <div style="color: #606266;font-size: 16px;font-weight: 600;">
                        <span>成果形式</span>
                        <a-tooltip placement="right">
                            <template #title>
                                <span>成果形式</span>
                            </template>
                            <question-circle-two-tone style="margin-left: 5px" />
                        </a-tooltip>
                    </div>
                    <div id="Output" style="width: 100%; height: 360px;"></div>
                </a-card>
            </a-col>
        </a-row>
        <a-row :gutter="16">
            <a-col :span="24">
                <a-card class="card" style="height: 35vh;margin-top: 20px;">
                <div style="color: #606266;font-size: 16px;font-weight: 600;">
                    <span>项目完成进度</span>
                    <a-tooltip placement="right">
                        <template #title>
                            <span>项目完成进度，(0-1)</span>
                        </template>
                        <question-circle-two-tone style="margin-left: 5px" />
                    </a-tooltip>
                </div>
                <div id="ProjectProgress" style="width: 100%; height: 360px;"></div>
            </a-card>
            </a-col>
            
        </a-row>
    </div>
</template>

<script setup>
import { QuestionCircleTwoTone } from '@ant-design/icons-vue'
import * as echarts from "echarts";
import { reactive, ref, onBeforeMount } from 'vue';
import { getSummary } from "../api/summary";
import { getSubscribeInfo } from '../api/subscribe';
import { useRouter } from 'vue-router'

const daysRange = ref(7);

const router = useRouter()

const data = reactive({
    Students: 0,
    projects: 0,
    // projectAmount: 0.00,
    outputs: 0,
})

onBeforeMount(() => {
    // subscribeInfo();
    initChart();
});

const initChart = () => {
    let param = {
        daysRange: daysRange.value
    }
    getSummary(param).then((res) => {
        if (res.data.code == 0) {
            data.Students = res.data.data.student_num
            data.projects = res.data.data.project_num
            data.outputs = res.data.data.output_num

            echarts.init(document.getElementById("Student")).setOption({
                tooltip: {
                    trigger: 'item'
                },
                legend: {
                    top: 'bottom',
                    left: 'center',
                },
                series: [
                    {
                        type: 'pie',
                        bottom: '15%',
                        radius: ['40%', '70%'],
                        avoidLabelOverlap: false,
                        itemStyle: {
                            borderRadius: 10,
                            borderColor: '#fff',
                            borderWidth: 2,
                        },
                        label: {
                            show: false,
                            position: 'center'
                        },
                        emphasis: {
                            label: {
                                show: true,
                                fontSize: 25,
                                fontWeight: 'bold'
                            }
                        },
                        labelLine: {
                            show: false
                        },
                        data: res.data.data.student_pan,
                    }
                ]
            })

            echarts.init(document.getElementById("Project")).setOption({
                tooltip: {
                    trigger: 'item'
                },
                legend: {
                    top: 'bottom',
                    left: 'center',
                },
                series: [
                    {
                        type: 'pie',
                        bottom: '15%',
                        radius: ['40%', '70%'],
                        avoidLabelOverlap: false,
                        itemStyle: {
                            borderRadius: 10,
                            borderColor: '#fff',
                            borderWidth: 2,
                        },
                        label: {
                            show: false,
                            position: 'center'
                        },
                        emphasis: {
                            label: {
                                show: true,
                                fontSize: 25,
                                fontWeight: 'bold'
                            }
                        },
                        labelLine: {
                            show: false
                        },
                        data: res.data.data.project_pan,
                    }
                ]
            })

            echarts.init(document.getElementById("Output")).setOption({
                tooltip: {
                    trigger: 'item'
                },
                legend: {
                    top: 'bottom',
                    left: 'center',
                },
                series: [
                    {
                        type: 'pie',
                        bottom: '15%',
                        radius: ['40%', '70%'],
                        avoidLabelOverlap: false,
                        itemStyle: {
                            borderRadius: 10,
                            borderColor: '#fff',
                            borderWidth: 2,
                        },
                        label: {
                            show: false,
                            position: 'center'
                        },
                        emphasis: {
                            label: {
                                show: true,
                                fontSize: 25,
                                fontWeight: 'bold'
                            }
                        },
                        labelLine: {
                            show: false
                        },
                        data: res.data.data.output_pan,
                    }
                ]
            })

            var myChart = echarts.init(document.getElementById('ProjectProgress'));
            var option;

            option = {
                color: ['#ee6666'],
                tooltip: {
                    trigger: 'axis',
                    axisPointer: {
                        type: 'shadow'
                    }
                },
                grid: {
                    left: '3%',
                    right: '4%',
                    bottom: '3%',
                    containLabel: true
                },
                xAxis: [
                    {
                        type: 'category',
                        data: res.data.data.progress_bar.map(obj => {return obj.name}),
                        axisTick: {
                            alignWithLabel: true
                        }
                    }
                ],
                yAxis: [
                    {
                        type: 'value'
                    }
                ],
                series: [
                    {
                        name: 'Direct',
                        type: 'bar',
                        barWidth: '60%',
                        data: res.data.data.progress_bar.map(obj => {return obj.value})
                    }
                ]
            };

            option && myChart.setOption(option);
        }
    })
}


// // 获取用户订阅信息
// const subscribeInfo = () => {
//     getSubscribeInfo().then((res) => {
//         if (res.data.code == 0 && res.data.data.version == 1) {
//             router.push('/result')
//         }
//     })
// }
</script>

<style scoped>
.card {
    border: none;
    box-shadow: 0 1px 16px 0 rgb(33 41 48 / 5%);
}
</style>