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

// ✅ LOGOUT CORRECTO
const logout = () => {
  localStorage.removeItem('token')
  localStorage.removeItem('user')
  router.push('/api/login')
}

const cargarProveedores = async () => {
  try {
    const token = localStorage.getItem('token')

    const res = await axios.get('http://localhost:8080/api/suppliers', {
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

    <div class="table-container" v-if="!loading && suppliersFiltrados.length">
      <table>
        <thead>
          <tr>
            <th>NIT</th>
            <th>Nombre</th>
            <th>Apellido</th>
            <th>Cédula</th>
            <th>Tipo Proveedor</th>
            <th>Tipo Persona</th>
            <th>Beneficiarios / Socios</th>
            <th>Datos Bancarios</th>
            <th>Estado</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="supplier in suppliersFiltrados" :key="supplier.id">
            <td>{{ supplier.nit }}</td>
            <td>{{ supplier.nombre }}</td>
            <td>{{ supplier.apellido }}</td>
            <td>{{ supplier.cedula }}</td>
            <td>{{ supplier.tipo_proveedor }}</td>
            <td>{{ supplier.tipo_persona }}</td>

            <td>
              <div v-if="supplier.beneficiarios?.length">
                <div v-for="(b, index) in supplier.beneficiarios" :key="index">
                  {{ b }}
                </div>
              </div>
              <span v-else>No registra</span>
            </td>

            <td>
              <div v-if="supplier.datos_bancarios?.length">
                <div v-for="(d, index) in supplier.datos_bancarios" :key="index">
                  {{ d }}
                </div>
              </div>
              <span v-else>No registra</span>
            </td>

            <td>
              <span :class="['estado', supplier.estado?.toLowerCase()]">
                {{ supplier.estado }}
              </span>
            </td>
          </tr>
        </tbody>
      </table>
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

/* ================= LOGOUT ================= */
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

.table-container {
  overflow-x: auto;
  background: rgba(255, 255, 255, 0.05);
  border-radius: 15px;
  padding: 15px;
}

table {
  width: 100%;
  border-collapse: collapse;
  min-width: 1200px;
}

thead {
  background: #0f172a;
}

th, td {
  padding: 12px;
  border-bottom: 1px solid #334155;
  vertical-align: top;
}

th {
  color: #38bdf8;
}


</style>
