// ── Scene Setup ────────────────────────────────────────────────
const container = document.getElementById('scene-container');
const W = container.clientWidth, H = container.clientHeight;

const scene    = new THREE.Scene();
const camera   = new THREE.PerspectiveCamera(50, W / H, 0.1, 100);
camera.position.set(3, 2.5, 4);
camera.lookAt(0, 0, 0);

const renderer = new THREE.WebGLRenderer({ antialias: true, alpha: true });
renderer.setSize(W, H);
renderer.setPixelRatio(window.devicePixelRatio);
renderer.setClearColor(0x000000, 0);
container.appendChild(renderer.domElement);

// ── Lighting ───────────────────────────────────────────────────
scene.add(new THREE.AmbientLight(0x6060ff, 0.4));

const dirLight = new THREE.DirectionalLight(0xffffff, 1.2);
dirLight.position.set(5, 8, 5);
scene.add(dirLight);

const fillLight = new THREE.DirectionalLight(0x4040ff, 0.5);
fillLight.position.set(-5, -3, -5);
scene.add(fillLight);

const rimLight = new THREE.PointLight(0x00ffcc, 1.0, 20);
rimLight.position.set(-4, 4, -4);
scene.add(rimLight);

// ── Cube ───────────────────────────────────────────────────────
const geo = new THREE.BoxGeometry(1.8, 1.8, 1.8);

// Six faces, six colors
const faceColors = [0x6655ff, 0x4499ff, 0x44ddaa, 0xff6644, 0xffcc44, 0xff44aa];
const materials  = faceColors.map(c =>
	new THREE.MeshPhongMaterial({ color: c, shininess: 80, specular: 0x222266 })
);
const wireMaterials = faceColors.map(c =>
	new THREE.MeshPhongMaterial({ color: c, shininess: 80, wireframe: true })
);

const cube = new THREE.Mesh(geo, materials);
scene.add(cube);

// Edges overlay
const edgesGeo = new THREE.EdgesGeometry(geo);
const edgesMat = new THREE.LineBasicMaterial({ color: 0xaaaaff, linewidth: 2 });
const edges = new THREE.LineSegments(edgesGeo, edgesMat);
cube.add(edges);

// ── Grid ───────────────────────────────────────────────────────
const grid = new THREE.GridHelper(10, 20, 0x222244, 0x111133);
grid.position.y = -2;
scene.add(grid);

// ── Orbit (manual, no OrbitControls dependency) ────────────────
let isDragging = false, lastX = 0, lastY = 0;
let rotX = 0.3, rotY = 0.5;
let autoRotate = true;
let scale = 1;

const canvas = renderer.domElement;

canvas.addEventListener('mousedown', e => { isDragging = true; lastX = e.clientX; lastY = e.clientY; });
window.addEventListener('mouseup',   () => { isDragging = false; });
window.addEventListener('mousemove', e => {
	if (!isDragging) return;
	rotY += (e.clientX - lastX) * 0.01;
	rotX += (e.clientY - lastY) * 0.01;
	lastX = e.clientX; lastY = e.clientY;
});

canvas.addEventListener('wheel', e => {
	camera.position.multiplyScalar(1 + e.deltaY * 0.001);
});

// Touch support
canvas.addEventListener('touchstart', e => {
	isDragging = true;
	lastX = e.touches[0].clientX; lastY = e.touches[0].clientY;
});
canvas.addEventListener('touchend', () => isDragging = false);
canvas.addEventListener('touchmove', e => {
	if (!isDragging) return;
	rotY += (e.touches[0].clientX - lastX) * 0.01;
	rotX += (e.touches[0].clientY - lastY) * 0.01;
	lastX = e.touches[0].clientX; lastY = e.touches[0].clientY;
});

// ── Controls ───────────────────────────────────────────────────
let wireframe = false, paused = false, exploded = false;
let explodeT = 0, explodeDir = 0;

// Compute face centre offsets for explode effect
const faceOffsets = [
	new THREE.Vector3( 1, 0, 0),
	new THREE.Vector3(-1, 0, 0),
	new THREE.Vector3( 0, 1, 0),
	new THREE.Vector3( 0,-1, 0),
	new THREE.Vector3( 0, 0, 1),
	new THREE.Vector3( 0, 0,-1),
];

// We'll fake explode by scaling cube non-uniformly via child meshes later;
// simpler: just scale the whole cube
document.getElementById('btn-wireframe').addEventListener('click', () => {
	wireframe = true;
	cube.material = wireMaterials;
	setActive('btn-wireframe');
});
document.getElementById('btn-solid').addEventListener('click', () => {
	wireframe = false;
	cube.material = materials;
	setActive('btn-solid');
});
document.getElementById('btn-pause').addEventListener('click', e => {
	paused = !paused;
	e.target.textContent = paused ? 'Resume' : 'Pause';
	e.target.classList.toggle('active', paused);
});
document.getElementById('btn-explode').addEventListener('click', e => {
	exploded = !exploded;
	explodeDir = exploded ? 1 : -1;
	e.target.classList.toggle('active', exploded);
});

function setActive(id) {
	['btn-wireframe','btn-solid'].forEach(b =>
	document.getElementById(b).classList.toggle('active', b === id)
	);
}

// ── Animation Loop ─────────────────────────────────────────────
let t = 0;
function animate() {
	requestAnimationFrame(animate);

	if (!paused) {
	t += 0.016;
	if (autoRotate) { rotY += 0.005; rotX += 0.002; }
	}

	// Explode: oscillate scale
	if (explodeDir !== 0) {
	explodeT = Math.max(0, Math.min(1, explodeT + explodeDir * 0.04));
	const s = 1 + explodeT * 0.6;
	cube.scale.set(s, s, s);
	if (explodeT <= 0 || explodeT >= 1) explodeDir = 0;
	}

	// Bob + rotation
	cube.position.y = Math.sin(t * 0.8) * 0.15;
	cube.rotation.x = rotX;
	cube.rotation.y = rotY;

	// Rim light orbit
	rimLight.position.x = Math.cos(t * 0.4) * 5;
	rimLight.position.z = Math.sin(t * 0.4) * 5;

	renderer.render(scene, camera);
}

animate();

// ── Resize ─────────────────────────────────────────────────────
window.addEventListener('resize', () => {
	const w = container.clientWidth, h = container.clientHeight;
	camera.aspect = w / h;
	camera.updateProjectionMatrix();
	renderer.setSize(w, h);
});
