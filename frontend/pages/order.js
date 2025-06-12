import React, { useEffect, useState } from 'react';
import axios from 'axios';
import Layout from '../components/Layout';
import OrderForm from '../components/OrderForm';

const OrderPage = () => {
  const [token, setToken] = useState('');

  useEffect(() => {
    const getToken = async () => {
      try {
        const res = await axios.get('http://localhost:8080/generate-token');
        setToken(res.data.token);
      } catch (err) {
        console.error(err);
      }
    };

    getToken();
  }, []);

  if (!token) return <div>Generando token...</div>;

  return (
    <Layout>
      <h1>Crear Orden</h1>
      <OrderForm token={token} />
    </Layout>
  );
};

export default OrderPage;
