<template>
    <div class="app-container">
        <div ref="mapContainer" class="map-container"></div>
        <canvas ref="canvasRef" class="canvas-container"></canvas>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import AMapLoader from '@amap/amap-jsapi-loader';
import { paths } from './mock';
import { latLngToMercator, mercatorToScreen } from '../utils/CoordUtils';


const mapContainer = ref<HTMLDivElement | null>(null);
let map: any = null;

let points: number[][] = [];
let point: number[] = [];
const canvasRef = ref<HTMLCanvasElement | null>(null)
let ctx: CanvasRenderingContext2D | null = null;

// 初始化地图
const initMap = (): void => {
    AMapLoader.load({
        key: 'e71234c84e3237b5a70ed3e1063d332f', // 重要：请替换为您从高德开放平台获取的API密钥
        version: '2.0',
        plugins: ['AMap.Scale', 'AMap.ToolBar', 'AMap.ControlBar']
    }).then((AMap: any) => {
        if (mapContainer.value) {
            if (canvasRef.value) {
                ctx = canvasRef.value.getContext('2d');
                canvasRef.value.width = mapContainer.value.offsetWidth;
                canvasRef.value.height = mapContainer.value.offsetHeight;
            }
            map = new AMap.Map(mapContainer.value, {
                viewMode: '3D',
                zoom: 11,
                center: [115.54085802537803, 39.16847619732222]
            });
            map.on('complete', () => {
                centerMercator = latLngToMercator(map.getCenter().lng, map.getCenter().lat);
                currentZoom = map.getZoom();
                drawRoute(paths);
            })
            map.on('moveend', () => {
                centerMercator = latLngToMercator(map.getCenter().lng, map.getCenter().lat);
                currentZoom = map.getZoom();
                drawRoute(paths);
            })

            // 添加控件
            map.addControl(new AMap.Scale());
            map.addControl(new AMap.ToolBar());
            map.addControl(new AMap.ControlBar({
                position: {
                    right: '10px',
                    top: '10px'
                }
            }));
        }
    }).catch((e: Error) => {
        console.error('高德地图加载失败：', e);
    });
};

let centerMercator: { x: number, y: number } = { x: 0, y: 0 };
let currentZoom: number = 0;

// 绘制路线
const drawRoute = (points: number[][]) => {
    if (!ctx) return;
    ctx.clearRect(0, 0, canvasRef.value?.width || 0, canvasRef.value?.height || 0);
    ctx.beginPath();
    ctx.strokeStyle = 'blue'; // 路线颜色
    ctx.lineWidth = 4;     // 路线宽度
    ctx.lineJoin = 'round'; // 线条连接处样式
    ctx.lineCap = 'round';  // 线条端点样式

    let firstPoint = points[0];
    let mercatorP = latLngToMercator(firstPoint[0], firstPoint[1]);
    let screenP = mercatorToScreen(mercatorP, centerMercator, currentZoom, 256, canvasRef.value?.width || 0, canvasRef.value?.height || 0);
    ctx.moveTo(screenP.x, screenP.y); // 移动到第一个点

    for (let i = 1; i < points.length; i++) {
        let nextPoint = points[i];
        let nextMercatorP = latLngToMercator(nextPoint[0], nextPoint[1]);
        let nextScreenP = mercatorToScreen(nextMercatorP, centerMercator, currentZoom, 256, canvasRef.value?.width || 0, canvasRef.value?.height || 0);
        ctx.lineTo(nextScreenP.x, nextScreenP.y); // 连接到下一个点
    }

    ctx.stroke(); // 绘制路径
}

onMounted(() => {
    initMap();
});
</script>

<style scoped lang="scss">
.app-container {
    width: 100%;
    height: 100%;
}

.map-container {
    width: 100%;
    height: 100%;
}

.canvas-container {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    pointer-events: none;
    z-index: 10;
}
</style>