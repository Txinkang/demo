// 假设地球半径 R = 6378137 米 (Web Mercator 标准)
const R = 6378137; // 地球半径，单位米
const MAX_LATITUDE = 85.05112878; // Web Mercator 投影的有效纬度范围

// 辅助函数：角度转弧度
function toRadians(degrees: number): number {
    return degrees * Math.PI / 180;
}

/**
 * 将经纬度转换为 Web Mercator 坐标 (x, y)，单位米
 * @param {number} lng 经度
 * @param {number} lat 纬度
 * @returns {{x: number, y: number}} Web Mercator 坐标
 */
export function latLngToMercator(lng: number, lat: number) {
    let x = R * toRadians(lng);
    let y = R * Math.log(Math.tan(Math.PI / 4 + toRadians(lat) / 2));

    // 限制纬度范围，避免无穷大
    if (y > R * Math.PI) y = R * Math.PI;
    if (y < -R * Math.PI) y = -R * Math.PI;

    return { x, y };
}

/**
 * 将 Web Mercator 坐标 (x, y) 转换为屏幕像素坐标
 * @param {{x: number, y: number}} mercatorCoords Web Mercator 坐标 (米)
 * @param {{x: number, y: number}} centerMercatorCoords 地图中心点的 Web Mercator 坐标 (米)
 * @param {number} zoom 缩放级别
 * @param {number} tileSize 瓦片尺寸 (默认为 256 或 512)
 * @param {number} mapWidth 地图容器宽度 (像素)
 * @param {number} mapHeight 地图容器高度 (像素)
 * @returns {{x: number, y: number}} 屏幕像素坐标
 */
export function mercatorToScreen(mercatorCoords: { x: number, y: number }, centerMercatorCoords: { x: number, y: number }, zoom: number, tileSize = 256, mapWidth: number, mapHeight: number) {
    // 每个瓦片代表的米数
    const resolution = (2 * Math.PI * R) / (tileSize * Math.pow(2, zoom));

    // 瓦片像素坐标 (相对于地图原点，通常是(0,0)在左上角)
    let pixelX = mercatorCoords.x / resolution;
    let pixelY = -mercatorCoords.y / resolution; // Y轴反向，因为屏幕坐标Y轴向下

    let centerPixelX = centerMercatorCoords.x / resolution;
    let centerPixelY = -centerMercatorCoords.y / resolution;

    // 计算相对于屏幕中心点的偏移
    let offsetX = pixelX - centerPixelX;
    let offsetY = pixelY - centerPixelY;

    // 转换为屏幕坐标 (以容器左上角为(0,0))
    let screenX = mapWidth / 2 + offsetX;
    let screenY = mapHeight / 2 + offsetY;

    return { x: screenX, y: screenY };
}