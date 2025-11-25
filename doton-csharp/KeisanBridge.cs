using System;

namespace KeisanTrading.Bridge
{
    /// <summary>
    /// Doton (Tierra) - Puente para NinjaTrader y ejecución de señales
    /// </summary>
    public class KeisanBridge
    {
        /// <summary>
        /// Ping - Validación de conexión
        /// </summary>
        public static string Ping()
        {
            return "Doton OK";
        }

        public static void Main(string[] args)
        {
            Console.WriteLine("Doton activo");
            string response = Ping();
            Console.WriteLine($"Ping response: {response}");
        }
    }
}
