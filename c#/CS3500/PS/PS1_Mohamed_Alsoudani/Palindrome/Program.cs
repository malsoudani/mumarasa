using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using BufferOptions;

namespace Palindrome
{
    class Program
    {
        /// <summary>
        /// Returns the words if they are Palindromes or returns nothing
        /// </summary>
        static void Main(string[] args)
        {
            string line;
            while ((line = Console.ReadLine()) != null)
            {
                foreach (string output in DisallowOutputBuffer.FilterIterator(line, s => s.Reverse().SequenceEqual(s) ? s : ""))
                {
                    Console.WriteLine(output);
                }
            }
        }
    }
}
