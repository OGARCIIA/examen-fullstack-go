import React, { useEffect, useState } from 'react';
import axios from 'axios';
import {
  TextField, Button, MenuItem, CircularProgress
} from '@mui/material';

const OrderForm = ({ token }) => {
  const [products, setProducts] = useState([]);
  const [selectedProduct, setSelectedProduct] = useState('');
  const [quantity, setQuantity] = useState(1);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    const fetchProducts = async () => {
      try {
        const res = await axios.get('http://localhost:8080/api/products', {
          headers: {
            Authorization: `Bearer ${token}`
          }
        });
        setProducts(res.data);
      } catch (err) {
        console.error(err);
      }
    };

    fetchProducts();
  }, [token]);

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);

    try {
      await axios.post('http://localhost:8080/api/orders', {
        product_id: selectedProduct,
        quantity: quantity
      }, {
        headers: {
          Authorization: `Bearer ${token}`
        }
      });

      alert('Orden creada con éxito!');
    } catch (err) {
      console.error(err);
      alert('Error al crear la orden');
    } finally {
      setLoading(false);
    }
  };

  if (products.length === 0) return <CircularProgress />;

  return (
    <form onSubmit={handleSubmit}>
      <TextField
        select
        label="Producto"
        value={selectedProduct}
        onChange={(e) => setSelectedProduct(e.target.value)}
        fullWidth
        margin="normal"
        required
      >
        {products.map((p) => (
          <MenuItem key={p.id} value={p.id}>
            {p.name} (Stock: {p.stock})
          </MenuItem>
        ))}
      </TextField>

      <TextField
        type="number"
        label="Cantidad"
        value={quantity}
        onChange={(e) => setQuantity(parseInt(e.target.value))}
        fullWidth
        margin="normal"
        required
        inputProps={{ min: 1 }}
      />

      <Button type="submit" variant="contained" color="primary" disabled={loading}>
        {loading ? 'Procesando...' : 'Crear Orden'}
      </Button>
    </form>
  );
};

export default OrderForm;
