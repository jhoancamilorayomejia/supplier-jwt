<template>
  <div class="dashboard">
    <h1>📋 Lista de Proveedores</h1>

    <!-- BOTÓN SALIR -->
    <div class="logout">
      <button @click="logout">Salir</button>
    </div>

    <!-- ERROR -->
    <div v-if="error" class="error">{{ error }}</div>

    <!-- TABLA -->
    <div class="table-container" v-if="!loading && suppliers.length">
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
            <th>Acciones</th>
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

            <!-- BENEFICIARIOS -->
            <td>
              <div v-if="supplier.beneficiarios && supplier.beneficiarios.length">
                <div v-for="(b, index) in supplier.beneficiarios" :key="index">
                  {{ b }}
                </div>
              </div>
              <span v-else>No registra</span>
            </td>

            <!-- DATOS BANCARIOS -->
            <td>
              <div v-if="supplier.datos_bancarios && supplier.datos_bancarios.length">
                <div v-for="(d, index) in supplier.datos_bancarios" :key="index">
                  {{ d }}
                </div>
              </div>
              <span v-else>No registra</span>
            </td>

            <!-- ACCIONES -->
            <td>
              <button class="btn-edit" @click="editarProveedor(supplier)">Editar</button>
              <button class="btn-delete" @click="eliminarProveedor(supplier.id)">Eliminar</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <p v-if="loading">Cargando proveedores...</p>

    <!-- BOTÓN AGREGAR -->
    <div class="add-btn-container">
      <button class="btn-add" @click="agregarProveedor">➕ Agregar Proveedor</button>
    </div>

    <!-- MODAL -->
    <div v-if="mostrarModal" class="modal-overlay" @click.self="cerrarModal">
      <div class="modal">
        <h2>{{ modoEdicion ? 'Editar Proveedor' : 'Agregar Proveedor' }}</h2>

        <form @submit.prevent="guardarProveedor">

          <input v-model="form.nit" placeholder="NIT" required />
          <input v-model="form.nombre" placeholder="Nombre" required />
          <input v-model="form.apellido" placeholder="Apellido" required />
          <input v-model="form.cedula" placeholder="Cédula" required />
          <input v-model="form.tipo_proveedor" placeholder="Tipo Proveedor" />
          <input v-model="form.tipo_persona" placeholder="Tipo Persona" />

          <!-- BENEFICIARIOS -->
          <label>Beneficiarios / Socios (uno por línea)</label>
          <textarea
            v-model="form.beneficiariosTexto"
            rows="4"
            placeholder="Juan Perez - 12345"
          ></textarea>

          <!-- DATOS BANCARIOS -->
          <label>Datos Bancarios (uno por línea)</label>
          <textarea
            v-model="form.datosBancariosTexto"
            rows="4"
            placeholder="Banco: Bancolombia | Cuenta: 123456"
          ></textarea>

          <div class="modal-actions">
            <button type="submit" class="btn-add">
              {{ modoEdicion ? 'Actualizar' : 'Guardar' }}
            </button>
            <button type="button" class="btn-delete" @click="cerrarModal">
              Cancelar
            </button>
          </div>
        </form>

      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import axios from 'axios'

const router = useRouter()
const suppliers = ref([])
const loading = ref(true)
const error = ref('')
const mostrarModal = ref(false)
const modoEdicion = ref(false)

const form = ref({
  id: null,
  nit: '',
  nombre: '',
  apellido: '',
  cedula: '',
  tipo_proveedor: '',
  tipo_persona: '',
  beneficiariosTexto: '',
  datosBancariosTexto: ''
})

// LOGOUT
const logout = () => {
  localStorage.clear()
  router.push('/api/login')
}

// MODAL
const cerrarModal = () => mostrarModal.value = false

const agregarProveedor = () => {
  modoEdicion.value = false
  form.value = {
    id: null,
    nit: '',
    nombre: '',
    apellido: '',
    cedula: '',
    tipo_proveedor: '',
    tipo_persona: '',
    beneficiariosTexto: '',
    datosBancariosTexto: ''
  }
  mostrarModal.value = true
}

// EDITAR
const editarProveedor = (supplier) => {
  modoEdicion.value = true
  form.value = {
    ...supplier,
    beneficiariosTexto: supplier.beneficiarios?.join('\n') || '',
    datosBancariosTexto: supplier.datos_bancarios?.join('\n') || ''
  }
  mostrarModal.value = true
}

// GUARDAR / ACTUALIZAR
const guardarProveedor = async () => {
  const token = localStorage.getItem('token')

  const payload = {
    ...form.value,
    beneficiarios: JSON.stringify(form.value.beneficiariosTexto.split('\n').filter(Boolean)),
    datos_bancarios: JSON.stringify(form.value.datosBancariosTexto.split('\n').filter(Boolean))
  }

  if (modoEdicion.value) {
    await axios.put(`http://localhost:8080/api/suppliers/${form.value.id}`, payload, {
      headers: { Authorization: `Bearer ${token}` }
    })
  } else {
    await axios.post(`http://localhost:8080/api/suppliers`, payload, {
      headers: { Authorization: `Bearer ${token}` }
    })
  }

  cerrarModal()
  cargarProveedores()
}

// ELIMINAR
const eliminarProveedor = async (id) => {
  if (!confirm('¿Eliminar proveedor?')) return
  const token = localStorage.getItem('token')
  await axios.delete(`http://localhost:8080/api/suppliers/${id}`, {
    headers: { Authorization: `Bearer ${token}` }
  })
  cargarProveedores()
}

// CARGAR
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
  } catch  {
    error.value = 'Error cargando proveedores'
  } finally {
    loading.value = false
  }
}

onMounted(cargarProveedores)
</script>



<style scoped>
.dashboard {
  min-height: 100vh;
  background: linear-gradient(135deg, #1f2937, #111827);
  padding: 30px;
  color: white;
}

/* ================= TÍTULO ================= */
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

/* ================= ERROR ================= */
.error {
  background: #dc2626;
  padding: 12px;
  border-radius: 8px;
  margin-bottom: 15px;
  text-align: center;
}

/* ================= TABLA ================= */
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

/* Contenido interno beneficiarios y bancos */
td div {
  margin-bottom: 6px;
  font-size: 14px;
}

/* ================= BOTONES ================= */
.add-btn-container {
  margin-top: 30px;
  
}

.btn-add {
  background: linear-gradient(135deg, #22c55e, #16a34a);
  color: white;
  border: none;
  padding: 12px 20px;
  border-radius: 25px;
  font-weight: bold;
  cursor: pointer;
  transition: 0.3s;
}

.btn-add:hover {
  transform: scale(1.05);
  background: linear-gradient(135deg, #16a34a, #15803d);
}

.btn-edit {
  background: #3b82f6;
  border: none;
  color: white;
  padding: 6px 12px;
  border-radius: 8px;
  cursor: pointer;
  margin-right: 5px;
}

.btn-edit:hover {
  background: #2563eb;
}

.btn-delete {
  background: #ef4444;
  border: none;
  color: white;
  padding: 6px 12px;
  border-radius: 38px;
  cursor: pointer;
}

.btn-delete:hover {
  background: #dc2626;
}

/* ================= MODAL ================= */
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0,0,0,0.7);
  display: flex;
  justify-content: center;
  align-items: flex-start;
  padding-top: 60px;
  z-index: 999;
}

.modal {
  background: #1f2937;
  width: 500px;
  max-height: 85vh;
  overflow-y: auto;
  border-radius: 20px;
  padding: 25px;
  animation: slideDown 0.4s ease;
  box-shadow: 0 10px 40px rgba(0,0,0,0.6);
}

@keyframes slideDown {
  from { transform: translateY(-80px); opacity: 0; }
  to { transform: translateY(0); opacity: 1; }
}

/* ================= FORMULARIO ================= */
.modal h2 {
  text-align: center;
  margin-bottom: 15px;
}

.modal label {
  font-weight: bold;
  color: #93c5fd;
  margin: 10px 0 5px;
  display: block;
}

.modal input,
.modal textarea {
  width: 100%;
  padding: 10px;
  margin-bottom: 12px;
  border-radius: 10px;
  border: none;
  outline: none;
  background: #334155;
  color: white;
}

.modal textarea {
  resize: vertical;
  min-height: 80px;
}

/* ================= BOTONES MODAL ================= */
.modal-actions {
  display: flex;
  justify-content: space-between;
  gap: 10px;
  margin-top: 10px;
}

/* ================= SCROLL MODAL ================= */
.modal::-webkit-scrollbar {
  width: 6px;
}

.modal::-webkit-scrollbar-thumb {
  background: #475569;
  border-radius: 10px;
}
</style>
