using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using BufferOptions;
using System.Text.RegularExpressions;

namespace Format
{
    class Program
    {
        static void Main(string[] args)
        {
            Regex regex = new Regex(@"^\d+$");
            if (args.Length != 1 || !regex.IsMatch(args[0]))
            {
                Console.Error.WriteLine("Provide a positive integer for the number of columns");
                System.Environment.Exit(1);
            }

            string line;
            int columns = int.Parse(args[0]);
            int i = 0;

            while ((line = Console.ReadLine()) != null)
            {
                string[] lineTokens = line.Split(new[] { " ", "\t", "\r", "\n", "\r\n" }, StringSplitOptions.RemoveEmptyEntries);
   
                foreach (string token in lineTokens)
                {
                    foreach (string output in DisallowOutputBuffer.FilterIterator(token, s => s))
                    {
                        i++;
                        Console.Write("{0}{1} ", output, i);
                        if ((i % columns) == 0)
                        {
                            Console.Write("\n");
                        } 
                    }
                }
            }
        }
    }
}
