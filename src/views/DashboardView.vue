<template>
  <div class="dashboard">
    <h1>📋 Lista de Proveedores</h1>

    <div v-if="error" class="error">{{ error }}</div>

    <div class="table-container" v-if="!loading">
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
          </tr>
        </thead>
        <tbody>
          <tr v-for="supplier in suppliers" :key="supplier.id">
            <td>{{ supplier.nit }}</td>
            <td>{{ supplier.nombre }}</td>
            <td>{{ supplier.apellido }}</td>
            <td>{{ supplier.cedula }}</td>
            <td>{{ supplier.tipo_proveedor }}</td>
            <td>{{ supplier.tipo_persona }}</td>

            <td>
              <div v-if="supplier.beneficiarios.length">
                <div v-for="b in supplier.beneficiarios" :key="b.cedula">
                  {{ b.nombre }} - {{ b.cedula }}
                </div>
              </div>
            </td>

            <td>
              <div v-if="supplier.datos_bancarios.length">
                <div v-for="(d, index) in supplier.datos_bancarios" :key="d.num_cuenta">
                  Banco: {{ d.banco }}<br />
                  Cuenta: {{ d.num_cuenta }}<br />
                  Tipo: {{ d.tipo_cuenta }}
                  <hr v-if="index !== supplier.datos_bancarios.length - 1" />
                </div>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <p v-if="loading">Cargando proveedores...</p>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import axios from 'axios'

const suppliers = ref([])
const loading = ref(true)
const error = ref('')
const router = useRouter()

onMounted(async () => {
  const token = localStorage.getItem('token')

  // ✅ Si no hay token, vuelve al login
  if (!token) {
    router.push('/api/login')
    return
  }

  try {
    const res = await axios.get('http://localhost:8080/api/suppliers', {
      headers: {
        Authorization: `Bearer ${token}`
      }
    })

    suppliers.value = res.data.map(s => ({
      ...s,
      beneficiarios: typeof s.beneficiarios === 'string'
        ? JSON.parse(s.beneficiarios)
        : (s.beneficiarios || []),
      datos_bancarios: typeof s.datos_bancarios === 'string'
        ? JSON.parse(s.datos_bancarios)
        : (s.datos_bancarios || [])
    }))

  } catch (err) {
    if (err.response?.status === 401) {
      localStorage.removeItem('token')
      router.push('/api/login')
    } else {
      error.value = 'Error al cargar proveedores'
    }
  } finally {
    loading.value = false
  }
})
</script>

<style scoped>
.dashboard {
  min-height: 100vh;
  background: linear-gradient(135deg, #1f2937, #111827);
  padding: 30px;
  color: white;
}

h1 {
  text-align: center;
  margin-bottom: 20px;
}

.error {
  background: #dc2626;
  padding: 10px;
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
}

th {
  color: #38bdf8;
}
</style>
