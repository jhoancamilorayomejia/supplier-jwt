<script setup>
import { ref, onMounted, computed } from 'vue'
import { useRouter } from 'vue-router'
import axios from 'axios'

const router = useRouter()

// usuario logueado (NIT)
const user = ref(localStorage.getItem('user'))

const suppliers = ref([])
const loading = ref(true)
const error = ref('')

// ✅ Solo proveedores del usuario logueado
const suppliersFiltrados = computed(() => {
  return suppliers.value.filter(s => s.nit === user.value)
})

// ✅ LOGOUT
const logout = () => {
  localStorage.removeItem('token')
  localStorage.removeItem('user')
  router.push('/api/login')
}

const cargarProveedores = async () => {
  try {
    const token = localStorage.getItem('token')

    const res = await axios.get('http://localhost:3000/api/suppliers', {
      headers: { Authorization: `Bearer ${token}` }
    })

    suppliers.value = res.data.map(s => ({
      ...s,
      beneficiarios: JSON.parse(s.beneficiarios || '[]'),
      datos_bancarios: JSON.parse(s.datos_bancarios || '[]')
    }))
  } catch {
    error.value = 'Error cargando proveedores'
  } finally {
    loading.value = false
  }
}

onMounted(cargarProveedores)
</script>

<template>
  <div class="dashboard">
    <h1>📋 Proveedor asociado al usuario: {{ user }}</h1>

    <!-- BOTÓN SALIR -->
    <div class="logout">
      <button @click="logout">Salir</button>
    </div>

    <div v-if="error" class="error">{{ error }}</div>

    <!-- FACTURA -->
    <div class="factura-container" v-if="!loading && suppliersFiltrados.length">
      <div
        class="factura"
        v-for="supplier in suppliersFiltrados"
        :key="supplier.id"
      >
        <h2>📄 Información del Proveedor</h2>

        <div class="factura-grid">
          <div class="fila">
            <span class="label">NIT</span>
            <span class="valor">{{ supplier.nit }}</span>
          </div>

          <div class="fila">
            <span class="label">Nombre completo</span>
            <span class="valor">{{ supplier.nombre }} {{ supplier.apellido }}</span>
          </div>

          <div class="fila">
            <span class="label">Cédula</span>
            <span class="valor">{{ supplier.cedula }}</span>
          </div>

          <div class="fila">
            <span class="label">Tipo Proveedor</span>
            <span class="valor">{{ supplier.tipo_proveedor }}</span>
          </div>

          <div class="fila">
            <span class="label">Tipo Persona</span>
            <span class="valor">{{ supplier.tipo_persona }}</span>
          </div>
        </div>

        <div class="seccion">
          <h3>👥 Beneficiarios / Socios</h3>
          <ul v-if="supplier.beneficiarios?.length">
            <li v-for="(b, index) in supplier.beneficiarios" :key="index">
              {{ b }}
            </li>
          </ul>
          <p v-else>No registra</p>
        </div>

        <div class="seccion">
          <h3>🏦 Datos Bancarios</h3>
          <ul v-if="supplier.datos_bancarios?.length">
            <li v-for="(d, index) in supplier.datos_bancarios" :key="index">
              {{ d }}
            </li>
          </ul>
          <p v-else>No registra</p>
        </div>

        <div class="estado-box">
          Estado:
          <span :class="['estado', supplier.estado?.toLowerCase()]">
            {{ supplier.estado }}
          </span>
        </div>
      </div>
    </div>

    <p v-if="loading">Cargando proveedores...</p>
    <p v-if="!loading && !suppliersFiltrados.length">
      No hay proveedores asociados a este usuario.
    </p>
  </div>
</template>

<style scoped>
.dashboard {
  min-height: 100vh;
  background: linear-gradient(135deg, #1f2937, #111827);
  padding: 30px;
  color: white;
}

h1 {
  text-align: center;
  margin-bottom: 25px;
}

/* LOGOUT */
.logout {
  position: absolute;
  top: 20px;
  right: 20px;
}

.logout button {
  padding: 10px 20px;
  border: none;
  border-radius: 20px;
  background: black;
  color: white;
  font-weight: bold;
  cursor: pointer;
  transition: 0.3s;
}

.logout button:hover {
  background: #333;
}

.error {
  background: #dc2626;
  padding: 12px;
  border-radius: 8px;
  margin-bottom: 15px;
  text-align: center;
}

/* ===== FACTURA ===== */

.factura-container {
  display: flex;
  justify-content: center;
}

.factura {
  background: rgba(255, 255, 255, 0.06);
  border-radius: 20px;
  padding: 25px;
  max-width: 900px;
  width: 100%;
  box-shadow: 0 10px 25px rgba(0,0,0,0.4);
}

.factura h2 {
  text-align: center;
  margin-bottom: 20px;
  color: #38bdf8;
}

/* Dos columnas estilo factura */
.factura-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 15px 30px;
  margin-bottom: 20px;
}

.fila {
  display: flex;
  flex-direction: column;
  background: rgba(0,0,0,0.3);
  padding: 12px 15px;
  border-radius: 10px;
}

.label {
  font-size: 13px;
  color: #94a3b8;
}

.valor {
  font-size: 16px;
  font-weight: bold;
}

/* Secciones */
.seccion {
  margin-top: 25px;
}

.seccion h3 {
  margin-bottom: 10px;
  color: #60a5fa;
}

.seccion ul {
  padding-left: 20px;
}

.seccion li {
  margin-bottom: 5px;
}

/* Estado */
.estado-box {
  margin-top: 30px;
  text-align: center;
  font-size: 18px;
}

.estado {
  display: inline-block;
  padding: 6px 14px;
  border-radius: 20px;
  font-weight: bold;
  margin-left: 10px;
  background: #475569;
}
</style>
