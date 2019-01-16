using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using BufferOptions;

namespace Palindrome
{
    /// <summary>
    /// Palindrome is a program that takes input and determines if it is a Palindrome
    /// </summary>
    class Program
    {
        /// <summary>
        /// Outputs the words if they are Palindromes or Outputs nothing
        /// </summary>
        static void Main(string[] args)
        {
            string line;
            while ((line = Console.ReadLine()) != null)
            {
                foreach (string output in DisallowOutputBuffer.FilterIterator(line, s => s.Reverse().SequenceEqual(s) ? s : ""))
                {
                    if (!String.IsNullOrEmpty(output))
                    {
                        Console.WriteLine(output);
                    }
                }
            }
        }
    }
}
