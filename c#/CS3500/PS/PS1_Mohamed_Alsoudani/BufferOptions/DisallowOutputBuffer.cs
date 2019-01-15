using System;
using System.Collections.Generic;

namespace BufferOptions
{
    public static class DisallowOutputBuffer
    {
        public static IEnumerable<string> FilterIterator(string token, Func<string, string> filterLogic)
        {
            string output = filterLogic(token);
            yield return output;
        }
    }
}
